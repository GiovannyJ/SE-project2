
/*
 *updates users password
 confirms that pwd and confirm password match
 confirms that pwd and pwd in db don't match 
 updates pwd and then returns to login page
 */
export async function updatePwd(event){
    event.preventDefault();
    const formData = new FormData(document.getElementById('resetPwdForm'));
    const updatePwdURL = 'http://localhost:8080/account/update'
    
    const cur_username      = formData.get('username');
    const cur_pwd           = formData.get('pwd');
    const cur_confirmPwd    = formData.get('confirm_pwd');
    
    if (cur_username === 'guest'){
        alert('Cannot Reset Guest Password');
        return;
    }
    
    if (cur_pwd != cur_confirmPwd){
        alert('Passwords do not match');
        return;
    }
   
    var query = "username=" + cur_username;
    const accountData = await getAccount(query);
    console.log(accountData)
    
    
    if (cur_pwd === accountData[0].pwd){
        alert("Password already used");
        return;
    }
    
    const old_body = {
        id:		   accountData[0].id,
        fname:      accountData[0].fname,
        lname:      accountData[0].lname,
        fullname:   accountData[0].fullname,
        email:      accountData[0].email,
        pwd:        accountData[0].pwd,
        pnum:      accountData[0].pnum,
        username:   accountData[0].username,
        accesslevel: accountData[0].accesslevel,
    }

    const new_body = {
        id:		   accountData[0].id,
        fname:      accountData[0].fname,
        lname:      accountData[0].lname,
        fullname:   accountData[0].fullname,
        email:      accountData[0].email,
        pwd:        cur_pwd,
        pnum:      accountData[0].pnum,
        username:   accountData[0].username,
        accesslevel: accountData[0].accesslevel,
    }
    
    try{
        const requestBody = {
            old: old_body,
            new: new_body,
        };
        

        const response = await fetch(updatePwdURL, {
            method: 'PATCH',
            body: JSON.stringify(requestBody),
            headers: {
                'Content-Type': 'application/json',
            },
        });
        if(!response.ok){
            throw new Error(`HTTP error! Status: ${response.headers}`);
        }
        
        alert("Password Updated!")
        window.location.href = 'login.html'
    }catch (error){
        console.error(`Error during patch method:`, error.message)
    }
}

async function getAccount(param) {
    var url = 'http://localhost:8080/account';
    
    if (param){
      url = 'http://localhost:8080/account?' + param;
    }
  
    try {
      const response = await fetch(url);
  
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
  
      const accountData = await response.json();
      return accountData
    } catch (error) {
      console.error('Error during GET request:', error.message);
    }
  }


export async function updateAccessLevel(){}