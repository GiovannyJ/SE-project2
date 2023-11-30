package conn

import (
	s "API/models"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type account = s.Account
type posts = s.Posts
type comments = s.Comment
type commentfullcontext = s.CommentFullContext
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
func Accounts_GET(params map[string]interface{}) ([]account, error) {
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
        return nil, err
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
			return values, err
		}
		values = append(values, t)
	}

    if err = result.Err(); err != nil{
		return values, err
    }

    return values, nil
}

/*
*TESTED WORKING
Gets all posts from table
*/
func Posts_GET(params map[string]interface{}) ([]posts, error) {
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
        return nil, err
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
			return values, nil
		}
		values = append(values, t)
	}

    if err = result.Err(); err != nil{
		return values, err
    }
        
    return values, nil
}

/*
*TESTED WORKING
gets all posts from database in full context
joined POSTS with ACCOUNTS and IMAGES
*/
func PostsFullContext_GET(params map[string]interface{}) ([]fullcontextpost, error) {
    var sql strings.Builder
    sql.WriteString("SELECT p.*, a.fullname, a.username, a.accesslvl, i.imgname FROM POSTS p")
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
			&author.Accesslvl,
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


/*
*TESTED WORKING
gets all comments associated with postID and query strings
*/
func Comments_GET(params map[string]interface{}, postID int) ([]comments, error) {
    var sql strings.Builder
    sql.WriteString("SELECT * FROM COMMENTS")

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
			postIDQuery := fmt.Sprintf(" AND postID = %d", postID)
			sql.WriteString(postIDQuery)
		}else{
			postIDQuery := fmt.Sprintf(" WHERE postID = %d", postID)
			sql.WriteString(postIDQuery)
		}
		sql.WriteString(orderby)
    }else{
		postIDQuery := fmt.Sprintf(" WHERE postID = %d", postID)
		sql.WriteString(postIDQuery)
	}


    result, err := connect(sql.String())
    if err != nil {
        return nil, err
    }
    defer result.Close()

    var values []comments
    
    for result.Next(){
		var t comments
		if err := result.Scan(
				&t.ID, 
				&t.PostID,
				&t.AuthorID,
				&t.Content,
				&t.NumUp,
				&t.NumDown,
				&t.PostedDate,
				); err != nil{
			return values, nil
		}
		values = append(values, t)
	}
	
    if err = result.Err(); err != nil{
		return values, err
    }
        
    return values, nil
}


/*
*TESTED WORKING
gets all comments associated with postID in full context
*/
func CommentsFullContext_GET(params map[string]interface{}, postID int) ([]commentfullcontext, error) {
    var sql strings.Builder
	sql.WriteString("SELECT c.id, c.numUp as numUpComments, c.numDown as numDownComments, ")
	sql.WriteString("c.postedDate as commentPostedDate, c.content, p.id as postID, p.title,")
	sql.WriteString("p.descr, p.genre, p.numUp, p.numDown, p.postedDate as postPostedDate,")
	sql.WriteString("i.imgname, ")
	sql.WriteString("a.username as posterAuthor, a.email as posterEmail, a.id, a.accesslvl  ")
	sql.WriteString(",a2.username as commenterAuthor, a2.email as commenterEmail, a2.id, a2.accesslvl")
	sql.WriteString(" FROM COMMENTS c" )
	sql.WriteString(" INNER JOIN POSTS p ON c.postID = p.id")
	sql.WriteString(" INNER JOIN IMAGES i ON p.picID = i.id" )
	sql.WriteString(" INNER JOIN ACCOUNTS a ON a.id  = p.authorID" )
	sql.WriteString(" INNER JOIN ACCOUNTS a2 on a2.id = c.authorID")


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
			postIDQuery := fmt.Sprintf(" AND postID = %d", postID)
			sql.WriteString(postIDQuery)
		}else{
			postIDQuery := fmt.Sprintf(" WHERE postID = %d", postID)
			sql.WriteString(postIDQuery)
		}
		sql.WriteString(orderby)
    }else{
		postIDQuery := fmt.Sprintf(" WHERE postID = %d", postID)
		sql.WriteString(postIDQuery)
	}

	result, err := connect(sql.String())
    if err != nil {
        return nil, err
    }
    defer result.Close()

    var values []commentfullcontext

    for result.Next() {
        var t commentfullcontext
		var commentInfo comments
		var postAuthor account
		var commentAuthor account
		var i images
		var postInfo posts

        if err := result.Scan(
			&commentInfo.ID,
			&commentInfo.NumUp,
			&commentInfo.NumDown,
			&commentInfo.PostedDate,
			&commentInfo.Content,
			&postInfo.ID,
			&postInfo.Title,
			&postInfo.Descr,
			&postInfo.Genre,
			&postInfo.NumUp,
			&postInfo.NumDown,
			&postInfo.PostedDate,
			&i.ImgName,
			&postAuthor.Username,
			&postAuthor.Email,
			&postAuthor.ID,
			&postAuthor.Accesslvl,
			&commentAuthor.Username,
			&commentAuthor.Email,
			&commentAuthor.ID,
			&commentAuthor.Accesslvl,
        ); err != nil {
            return values, err
        }
		t.CommentInfo = commentInfo
		t.CommenterInfo = commentAuthor
		t.PostInfo = postInfo
		t.PostAuthorInfo = postAuthor
		t.ImageInfo = i
		t.CommentInfo.PostID = t.PostInfo.ID

        values = append(values, t)
    }

    if err = result.Err(); err != nil {
        return values, err
    }

    return values, nil
}

/*
*TESTED WORKING
gets all info from IMAGES table
*/
func Images_GET(params map[string]interface{}) ([]images, error) {
    var sql strings.Builder
    sql.WriteString("SELECT * FROM IMAGES")

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

    var values []images
    
    for result.Next(){
		var t images
		if err := result.Scan(
				&t.ID, 
				&t.ImgName, 
				&t.Size,
				&t.Date,
				); err != nil{
			return values, nil
		}
		values = append(values, t)
	}

    if err = result.Err(); err != nil{
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

    exists, err := Accounts_GET(exist_acc)
	if err != nil{
		return false
	}
    
    return len(exists) > 0
}

/*
*TESTED WORKING
POSTS TO POSTS table
DATA: model after Posts struct
RETURN: error when applicable nil when not
*/
func CreateNewPost(data posts) error {
	sql := fmt.Sprintf("INSERT INTO POSTS(title, descr, genre, authorID, picID, postedDate) VALUES('%s', '%s', '%s', '%d', '%d', CURDATE())",
		 data.Title, data.Descr, data.Genre, data.AuthorID, data.PicID)
	result, err := connect(sql)
	if err != nil {
	    return err
	}
	defer result.Close()

	return nil
}

/*
*TESTED WORKING
creates new comment when body supplied with postID, authorId, and content
*/
func CreateNewComment(data comments) error {
	sql := fmt.Sprintf("INSERT INTO COMMENTS(postID, authorID, content, postedDate) VALUES('%d', '%d', '%s', CURDATE())",
		 data.PostID, data.AuthorID, data.Content)
	result, err := connect(sql)
	if err != nil {
	    return err
	}
	defer result.Close()

	return nil
}

/*
*TESTED WORKING
uploads image information into the database
*/
func CreateNewImage(data images) error{
	sql := fmt.Sprintf("INSERT INTO IMAGES(imgname, size, date) VALUES('%s', '%s', CURDATE())",
			data.ImgName, data.Size)
	result, err := connect(sql)
	if err != nil{
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
        tableName   string
        setValues   []string
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
    case reflect.TypeOf(comments{}):
        tableName = "COMMENTS"
    }

    // Loop through fields of newType to construct SET part of the query
    for i := 0; i < newType.NumField(); i++ {
        field := newType.Field(i)
        columnName := strings.ToLower(field.Name)
        columnValue := newValue.Field(i).Interface()

        setValues = append(setValues, fmt.Sprintf("%s='%v'", columnName, columnValue))
    }

    // Only include id from oldType in WHERE clause
    for i := 0; i < oldType.NumField(); i++ {
        field := oldType.Field(i)
        columnName := strings.ToLower(field.Name)
        if columnName == "id" {
            columnValue := oldValue.Field(i).Interface()
            whereValues = append(whereValues, fmt.Sprintf("%s='%v'", columnName, columnValue))
        }
    }

    sql := fmt.Sprintf("UPDATE %s SET %s WHERE %s", tableName, strings.Join(setValues, ","), strings.Join(whereValues, " AND "))

    result, err := connect(sql)
    if err != nil {
        return err
    }
    defer result.Close()

    
    return nil
}

//**+++++++++++++++++++++DELETE QUERIES++++++++++++++++++++++++++++

/*
*TESTED WORKING
DELETES ACCOUNT from database
returns error if applicable
*/
func DeleteAccount(user int) error{
	// sql := fmt.Sprintf("DELETE a, p, c FROM ACCOUNTS a LEFT JOIN POSTS p ON a.id = p.authorID LEFT JOIN COMMENTS c ON a.id = c.authorID WHERE a.id = %d;", user)
	deleteCommentsFromPostsSQL := fmt.Sprintf("DELETE FROM COMMENTS WHERE postID IN (SELECT id FROM POSTS WHERE authorID = %d)", user)
	deleteCommentsSQL := fmt.Sprintf("DELETE FROM COMMENTS WHERE authorID = %d", user)
	deletePostsSQL := fmt.Sprintf("DELETE FROM POSTS WHERE authorID = %d", user)

	deleteAccountSQL := fmt.Sprintf("DELETE FROM ACCOUNTS WHERE id = %d", user)
	
	execute(deleteCommentsFromPostsSQL)
	execute(deleteCommentsSQL)
	execute(deletePostsSQL)
	execute(deleteAccountSQL)
	
	return nil
}


/*
*TESTED WORKING
DELETES POST from database
returns error if applicable
*/
func DeletePost(user int) error{
	deletePostSQL := fmt.Sprintf("DELETE FROM POSTS WHERE id=%d", user)
	deleteCommentsSQL := fmt.Sprintf("DELETE FROM COMMENTS WHERE postID=%d", user)

	execute(deleteCommentsSQL)
	execute(deletePostSQL)
	return nil
}

func execute(sql string) error{
	result, err := connect(sql)
	if err !=nil{
		return err
	}
	defer result.Close()
	return nil
}

/*
*TESTED WORKING
deletes comment from database
returns error if applicable
*/
func DeleteComment(commentID int) error{
	sql := fmt.Sprintf("DELETE FROM COMMENTS WHERE id= %d ", commentID)
	result, err := connect(sql)

	if err != nil{
		return err
	}
	defer result.Close()
	return nil
}