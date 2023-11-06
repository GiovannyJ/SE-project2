package conn


import (
	s "API/models"
	"database/sql"
	"fmt"
	"log"
	// "reflect"
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

//**MAKE ALL GENERIC DATABASE QUERIES HERE

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
				&t.Password,
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
/*
*TESTED WORKING
CREATE AN ACCOUNT
data: model after login struct

return error if applicable
*/
// func NewAccount(data login) error {
//     tableName := "ACCOUNT_INFO"

//     valueType := reflect.TypeOf(data)
//     value := reflect.ValueOf(data)

//     var columns []string
//     var values []string

//     if(accountExist(*data.Username)){
//         return &reflect.ValueError{}
//     }

//     for i := 0; i < valueType.NumField(); i++ {
//         field := valueType.Field(i)
//         columnName := strings.ToLower(field.Name)
//         columnValue := value.Field(i).Interface()

        
        
//         if columnName == "id" || columnName == "pwd" {
//             continue
//         }

//         if columnName == "password" {
//             if pwd, ok := columnValue.(*string); ok && pwd != nil {
//                 hashedPwd, err := HashPassword(*pwd)
//                 if err != nil {
//                     return err
//                 }
//                 columnValue = hashedPwd
//                 columnName = "pwd"
//             }
//         }

//         if ptr, ok := columnValue.(*string); ok && ptr != nil {
//             columnValue = *ptr
//         } else if ptr, ok := columnValue.(*int); ok && ptr != nil {
//             columnValue = *ptr
//         }

//         columns = append(columns, columnName)
//         values = append(values, fmt.Sprintf("'%v'", columnValue))
//     }

//     sql := fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s)", tableName, strings.Join(columns, ","), strings.Join(values, ","))
//     result, err := connect(sql)

//     if err != nil {
//         return err
//     }
//     defer result.Close()
// }
  /*
!NEEDS FIXING
*TESTED PASSING
UPDATING DATA ON THE health, pr_tracker, and account table
old data must grab all values and be shaped like table struct
new data must grab all values and be shaped like table struct
note: make sure that columns are not empty bc then it will not work

return bool- true if worked and false if not
*/
// func UpdateData(oldData interface{}, newData interface{}) error {
//     var (
//         tableName string
//         setValues []string
//         whereValues []string
//     )

    
//     oldType := reflect.TypeOf(oldData)
//     oldValue := reflect.ValueOf(oldData)
//     newType := reflect.TypeOf(newData)
//     newValue := reflect.ValueOf(newData)

//     switch oldType {
//         case reflect.TypeOf(account{}):
//             tableName = "ACCOUNT_INFO"
//         case reflect.TypeOf(healthdata{}):
//             tableName = "HEALTH_INFO"
//         case reflect.TypeOf(pr{}):
//             tableName = "PR_TRACKER"
//     }

//     for i := 0; i < newType.NumField(); i++ {
//         field := newType.Field(i)
//         columnName := strings.ToLower(field.Name)
//         columnValue := newValue.Field(i).Interface()

//         setValues = append(setValues, fmt.Sprintf("%s='%v'", columnName, columnValue))
//     }

//     for i := 0; i < oldType.NumField(); i++ {
//         field := oldType.Field(i)
//         columnName := strings.ToLower(field.Name)
//         columnValue := oldValue.Field(i).Interface()

//         whereValues = append(whereValues, fmt.Sprintf("%s='%v'", columnName, columnValue))
//     }

//     sql := fmt.Sprintf("UPDATE %s SET %s WHERE %s", tableName, strings.Join(setValues, ","), strings.Join(whereValues, " AND "))
    
//     result, err := connect(sql)
//     if err != nil {
//         return err
//     }
//     defer result.Close()

//     return nil
// }