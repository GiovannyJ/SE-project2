package conn

import (
	s "API/models"
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type account = s.Account
/*
*TEST PASSING
CONNECTION TO DATABASE
query: sql query to run
returns the sql values or error
*/
func connect(query string) (*sql.Rows, error) {
	username    :=     EnvVar("DB_USERNAME")
    password    :=     EnvVar("DB_PASSWORD")
    host        :=     EnvVar("DB_HOST")
    port        :=     EnvVar("DB_PORT")
    name        :=     EnvVar("DB_NAME")

    configOS := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, name)
    
    db, err := sql.Open("mysql", configOS)
	if err != nil{
        log.Fatal()
	}
    defer db.Close()

	res, err := db.Query(query)

	if err != nil{
		panic(err.Error())
	}
	
	return res, err
}

//**+++++++++++++++++++++SELECT QUERIES++++++++++++++++++++++++++++

/*
*TESTED WORKING
getting all accounts query
returns: all accounts in database
*/
func Accounts_GET(params map[string]interface{}) ([]account) {
    var sql strings.Builder
    sql.WriteString("SELECT * FROM ACCOUNT_INFO")

	if len(params) > 0 { 
        sql.WriteString(" WHERE ")
        var conditions []string
        for name, value := range params {
            condition := fmt.Sprintf("%s='%v'", name, value)
            conditions = append(conditions, condition)
        }

        sql.WriteString(strings.Join(conditions, " AND "))
    }


    result, err := connect(sql.String())
    if err != nil {
        return nil
    }
    defer result.Close()

    var values []account
    
    for result.Next(){
		var t account
		if err := result.Scan(
				&t.ID, 
				&t.Fname, 
				&t.Lname,
				&t.Fullname,
				&t.Email,
				&t.Pwd,
				&t.Pnum,
				&t.Age,
				&t.Username,
				); err != nil{
			return values
		}
		values = append(values, t)
	}

    if err = result.Err(); err != nil{
		return values
    }
        
    return values
}


//**+++++++++++++++++++++INSERT QUERIES++++++++++++++++++++++++++++

/*
*TESTED WORKING
POST TO: ACCOUNT_INFO table
DATA: model after account struct
RETURN: error when applicable nil if no error
*/
func CreateNewAccount(data account) error {
	if accountExist(data.Username) {
		return &s.AccountExistsError{Username: data.Username}
	}

	sql := fmt.Sprintf("INSERT INTO ACCOUNT_INFO(fname, lname, fullname, email, pwd, pnum, age, username) VALUES('%s', '%s', '%s', '%s', '%s', %d, %d, '%s')",
		data.Fname, data.Lname, data.Fullname, data.Email, data.Pwd, data.Pnum, data.Age, data.Username)

	result, err := connect(sql)
	if err != nil {
	    return err
	}
	defer result.Close()

	return nil
}

func accountExist(id string) bool{
    exist_acc := map[string]interface{}{"username": id}

    exists := Accounts_GET(exist_acc)
    
    return len(exists) > 0
}


/*
*TESTED PASSING
POST TO: health and tracker tables
data: model after struct for table
note: id must be empty since auto populates

return true if worked and false if not
*/
// func PostData(data interface{}) error {
//     var (
//         tableName string
//         columns   []string
//         values    []string
//     )

//     valueType := reflect.TypeOf(data)
//     value := reflect.ValueOf(data)

//     switch valueType {
//     case reflect.TypeOf(healthdata{}):
//         tableName = "HEALTH_INFO"
//     case reflect.TypeOf(pr{}):
//         tableName = "PR_TRACKER"
//     }

//     for i := 0; i < valueType.NumField(); i++ {
//         field := valueType.Field(i)
//         columnName := strings.ToLower(field.Name)
//         columnValue := value.Field(i).Interface()

//         // Dereference pointers if they are not nil
//         if ptr, ok := columnValue.(*string); ok && ptr != nil {
//             columnValue = *ptr
//         } else if ptr, ok := columnValue.(*int); ok && ptr != nil {
//             columnValue = *ptr
//         }

//         if columnName != "id"{
//             columns = append(columns, columnName)
//             values = append(values, fmt.Sprintf("'%v'", columnValue))
//         }
//     }

//     sql := fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s)", tableName, strings.Join(columns, ","), strings.Join(values, ","))

//     result, err := connect(sql)
//     if err != nil {
//         return err
//     }
//     defer result.Close()

//     return nil
// }

//**+++++++++++++++++++++UPDATE QUERIES++++++++++++++++++++++++++++



  /*
!NEEDS FIXING
*TESTED PASSING
GENERIC UPDATE METHOD WILL MATCH INTERFACE OF OLD AND NEW DATA
RETURN: error if applicable
!NEED TO IMPLEMENT WAY TO CHECK IF CHANGES ARE MADE PROPERLY
*/
func UpdateData(oldData interface{}, newData interface{}) error {
    var (
        tableName string
        setValues []string
        whereValues []string
    )

    
    oldType := reflect.TypeOf(oldData)
    oldValue := reflect.ValueOf(oldData)
    newType := reflect.TypeOf(newData)
    newValue := reflect.ValueOf(newData)

    switch oldType {
        case reflect.TypeOf(account{}):
            tableName = "ACCOUNT_INFO"
    }

    for i := 0; i < newType.NumField(); i++ {
        field := newType.Field(i)
        columnName := strings.ToLower(field.Name)
        columnValue := newValue.Field(i).Interface()

        setValues = append(setValues, fmt.Sprintf("%s='%v'", columnName, columnValue))
    }

    for i := 0; i < oldType.NumField(); i++ {
		field := oldType.Field(i)
		columnName := strings.ToLower(field.Name)
		columnValue := oldValue.Field(i).Interface()
	
		whereValues = append(whereValues, fmt.Sprintf("%s='%v'", columnName, columnValue))
	}

    sql := fmt.Sprintf("UPDATE %s SET %s WHERE %s", tableName, strings.Join(setValues, ","), strings.Join(whereValues, " AND "))
	
    result, err := connect(sql)
    if err != nil {
        return err
    }
    defer result.Close()
	
	if newType == reflect.TypeOf(account{}) {
		//check if new data is on database and old data is not on database
        if checkUpdate(newData.(account)) && !checkUpdate(oldData.(account)) {
            return nil
        } else {
            return &s.UpdateNotCompleteError{Msg: "Update Was not complete"}
        }
    }


    return nil
}
//*RETURN: true if update was made properly
func checkUpdate(newData account) bool{
	exist_acc := map[string]interface{}{
		"ID"       : newData.ID,
		"Fname"    : newData.Fname,
		"Lname"    : newData.Lname,
		"Fullname" : newData.Fullname,
		"Email"    : newData.Email,
		"Pwd" 	   : newData.Pwd,
		"Pnum"     : newData.Pnum,
		"Age"      : newData.Age,
		"Username" : newData.Username,
	}
	
	results := Accounts_GET(exist_acc)
	
	return len(results) > 0
}
