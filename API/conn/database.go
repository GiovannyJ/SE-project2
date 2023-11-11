package conn

import (
	s "API/models"
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

type account = s.Account
type posts = s.Posts
type fullcontextpost = s.FullContextPost
type images = s.Images

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
    sql.WriteString("SELECT * FROM ACCOUNTS")

	if len(params) > 0 {
		var conditions []string
		var orderby string
        for name, value := range params {
			if name == "order"{
				orderby = fmt.Sprintf(" ORDER BY %s", value)
				}else{
				condition := fmt.Sprintf("%s='%v'", name, value)
				conditions = append(conditions, condition)
			}
        }
		if len(conditions) > 0{
			sql.WriteString(" WHERE ")
			sql.WriteString(strings.Join(conditions, " AND "))
		}
		sql.WriteString(orderby)
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
			&t.Username,
			&t.Accesslvl,
		); err != nil {
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
*TESTED WORKING
Gets all posts from table
*/
func Posts_GET(params map[string]interface{}) ([]posts) {
    var sql strings.Builder
    sql.WriteString("SELECT * FROM POSTS")

	if len(params) > 0 {
		var conditions []string
		var orderby string
        for name, value := range params {
			if name == "order"{
				orderby = fmt.Sprintf(" ORDER BY %s", value)
				}else{
				condition := fmt.Sprintf("%s='%v'", name, value)
				conditions = append(conditions, condition)
			}
        }
		if len(conditions) > 0{
			sql.WriteString(" WHERE ")
			sql.WriteString(strings.Join(conditions, " AND "))
		}
		sql.WriteString(orderby)
    }
    result, err := connect(sql.String())
    if err != nil {
        return nil
    }
    defer result.Close()

    var values []posts
    
    for result.Next(){
		var t posts
		if err := result.Scan(
				&t.ID, 
				&t.Title, 
				&t.Descr,
				&t.Genre,
				&t.AuthorID,
				&t.NumUp,
				&t.NumDown,
				&t.PicID,
				&t.PostedDate,
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
*TESTED WORKING
gets all posts from database in full context
joined POSTS with ACCOUNTS and IMAGES
*/
func PostsFullContext_GET(params map[string]interface{}) ([]fullcontextpost, error) {
    var sql strings.Builder
    sql.WriteString("SELECT p.*, a.fullname, a.username, i.imgname FROM POSTS p")
    sql.WriteString(" INNER JOIN ACCOUNTS a ON p.authorID = a.id")
    sql.WriteString(" INNER JOIN IMAGES i ON p.picID = i.id")

    if len(params) > 0 {
		var conditions []string
		var orderby string
        for name, value := range params {
			if name == "order"{
				orderby = fmt.Sprintf(" ORDER BY %s", value)
				}else{
				condition := fmt.Sprintf("%s='%v'", name, value)
				conditions = append(conditions, condition)
			}
        }
		if len(conditions) > 0{
			sql.WriteString(" WHERE ")
			sql.WriteString(strings.Join(conditions, " AND "))
		}
		sql.WriteString(orderby)
    }
	result, err := connect(sql.String())
    if err != nil {
        return nil, err
    }
    defer result.Close()

    var values []fullcontextpost

    for result.Next() {
        var t fullcontextpost
        var author account
        var image images

        if err := result.Scan(
            &t.ID,
            &t.Title,
            &t.Descr,
            &t.Genre,
			&author.ID,
            &t.NumUp,
            &t.NumDown,
			&image.ID,
			&t.PostedDate,
            &author.Fullname,
            &author.Username,
			&image.ImgName,
        ); err != nil {
            return values, err
        }

        t.AuthorInfo = author
        t.ImageInfo = image
        values = append(values, t)
    }

    if err = result.Err(); err != nil {
        return values, err
    }

    return values, nil
}


//**+++++++++++++++++++++INSERT QUERIES++++++++++++++++++++++++++++

/*
*TESTED WORKING
POST TO: ACCOUNTS table
DATA: model after account struct
RETURN: error when applicable nil if no error
*/
func CreateNewAccount(data account) error {
	if accountExist(data.Username) {
		return &s.AccountExistsError{Username: data.Username}
	}

	sql := fmt.Sprintf("INSERT INTO ACCOUNTS(fname, lname, fullname, email, pwd, pnum, username, accesslvl) VALUES('%s', '%s', '%s', '%s', '%s', %d, '%s', 'user')",
		data.Fname, data.Lname, data.Fullname, data.Email, data.Pwd, data.Pnum, data.Username)
	result, err := connect(sql)
	if err != nil {
	    return err
	}
	defer result.Close()

	return nil
}
//*HELPER METHOD: checks if the account exits already RETURNS true if account exits and false if not
func accountExist(id string) bool{
    exist_acc := map[string]interface{}{"username": id}

    exists := Accounts_GET(exist_acc)
    
    return len(exists) > 0
}

/*
*TESTED WORKING
POSTS TO POSTS table
DATA: model after Posts struct
RETURN: error when applicable nil when not
*/
func CreateNewPost(data posts) error {
	sql := fmt.Sprintf("INSERT INTO POSTS(title, descr, genre, authorID, picID, postedDate) VALUES('%s', '%s', '%s', '%d', 1, CURDATE())",
		 data.Title, data.Descr, data.Genre, data.AuthorID)
	result, err := connect(sql)
	if err != nil {
	    return err
	}
	defer result.Close()

	return nil
}


/*
*TESTED WORKING
Using either email or username, compares password to value in database
RETURNS account if password valid nil if not
*/
func GetLoginInfo(username *string, password *string, email *string) (*account, error) {
    var sql strings.Builder
    q := "SELECT id, email, username, pwd FROM ACCOUNTS"

    switch {
    case username != nil:
        q += fmt.Sprintf(" WHERE username='%s'", *username)
    case email != nil:
        q += fmt.Sprintf(" WHERE email='%s'", *email)
    default:
        break
    }

    sql.WriteString(q)
    result, err := connect(sql.String())
    if err != nil {
        return nil, err
    }
    defer result.Close()

    if result.Next() {
        var t account
        if err := result.Scan(
            &t.ID,
            &t.Email,
            &t.Username,
			&t.Pwd,
        ); err != nil {
            return nil, err
        }
        if password != nil && t.Pwd != *password {
            return nil, errors.New("invalid password")
        }
        return &t, nil
    }
    return nil, errors.New("user not found")
}

/*
*TESTED WORKING
Grabs guest account from database and returns its values
*/
func GuestLogin() ([]account){
    q := "SELECT * FROM ACCOUNTS WHERE username = 'guest' and pwd='guest'"
	result, err := connect(q)
	if err != nil{
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
			&t.Username,
			&t.Accesslvl,
		); err != nil {
			return values
		}
		values = append(values, t)
	}

    if err = result.Err(); err != nil{
		return values
    }

    return values
}


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
            tableName = "ACCOUNTS"
		case reflect.TypeOf(posts{}):
			tableName = "POSTS"
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
		"Username" : newData.Username,
		"Accesslvl": newData.Accesslvl,
	}
	
	results := Accounts_GET(exist_acc)
	
	return len(results) > 0
}

//**+++++++++++++++++++++DELETE QUERIES++++++++++++++++++++++++++++

/*
*TESTED WORKING
DELETES ACCOUNT from database
returns error if applicable
*/
func DeleteAccount(user int) error{
	sql := fmt.Sprintf("DELETE FROM ACCOUNTS WHERE id=%d", user)
	result, err := connect(sql)
	
	if err != nil{
		return err
	}
	defer result.Close()
	return nil
}

/*
*TESTED WORKING
DELETES POST from database
returns error if applicable
*/
func DeletePost(user int) error{
	sql := fmt.Sprintf("DELETE FROM POSTS WHERE id=%d", user)
	result, err := connect(sql)
	
	if err != nil{
		return err
	}
	defer result.Close()
	return nil
}