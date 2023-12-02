/**
 * ==================================================================
 * DELETE METHODS
 * used to delete obj from database using html text input element
 * calls to delete endpoint in API
 * ==================================================================
 */

/**
 * deletes account using accountID
*/
export async function deleteAccount(accountID) {
  const userConfirmed = window.confirm("Are you sure you would like to delete this account?");
    if (!userConfirmed) {
        alert("Account was not deleted");
        return;
    } 
    
    const url = `http://localhost:8080/account/delete/${accountID}`;
  
    try {
      const response = await fetch(url, {
        method: 'DELETE',
      });
  
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
  
      alert('Account deleted successfully!');
      
      window.location.reload();

    } catch (error) {
      console.error('Error during DELETE request:', error.message);
    }
  }
  
  /**
   * deletes post using postID
  */
export async function deletePost(postID) {
    const userConfirmed = window.confirm("Are you sure you would like to delete this Post?");
    if (!userConfirmed) {
        alert("Post was not deleted");
        return;
    } 

    const url = `http://localhost:8080/posts/delete/${postID}`;
  
    try {
      const response = await fetch(url, {
        method: 'DELETE',
      });
  
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
  
      alert('Post deleted successfully!');
      window.location.reload()
    } catch (error) {
      console.error('Error during DELETE request:', error.message);
    }
  }
  
/**
 * Deletes comment from database using id
*/
export async function deleteComment(commentId) {
  const userConfirmed = window.confirm("Are you sure you would like to delete this Comment?");
    if (!userConfirmed) {
        alert("Comment was not deleted");
        return;
    } 
  const url = `http://localhost:8080/posts/delete/comment/${commentId}`;

  try {
    const response = await fetch(url, {
      method: 'DELETE',
    });

    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }

    alert('Comment deleted successfully!');
  } catch (error) {
    console.error('Error during DELETE request:', error.message);
  }
}

//*displays account(s) on the deleteAccountTable
export async function displayAccount(event, type) {
    event.preventDefault();
    var url = 'http://localhost:8080/account';
    var param = ""
    switch (type){
      case "delete":
        param = document.getElementById('accountSearchInput').value;
        break;
      case "changePerm":
        param = document.getElementById('promoteAcc').value;
        break;
    }

    if (param) {
        url = 'http://localhost:8080/account?id=' + param;
    }

    try {
        const response = await fetch(url);

        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
        }

        const accountData = await response.json();
        switch (type){
          case "delete":
            displayData('deleteAccountTable', accountData, "account");
            break;
          case "changePerm":
            displayData('promoteAccTable', accountData, "accountUpd");
            break;
        }

    } catch (error) {
        console.error('Error during GET request:', error.message);
    }
}

//*displays posts(s) on the deletePostsTable
export async function displayPost(event) {
  event.preventDefault();

  const postID = document.getElementById('postSearchInput').value
  
  var url = `http://localhost:8080/posts?id=${postID}`;
  

  try {
    const response = await fetch(url);

    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }

    const postsData = await response.json();
    console.log(postsData[0])
    
    displayData('deletePostTable', postsData, "post");
    
  } catch (error) {
    console.error('Error during GET request:', error.message);
  }
}

//*displays comments(s) on the deleteCommentTable
export async function displayComments(event) {
  event.preventDefault()
  const postId = document.getElementById('postID').value;
  if(!postId){
    alert("Please Enter a Post ID");
    return;
  }
  const param  = document.getElementById('commentSearchInput').value;
  var url = `http://localhost:8080/posts/${postId}/comments?`;

  if (param){
    url = url + "id="+param;
  }

  try {
    const response = await fetch(url);

    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }

    const commentsData = await response.json();
    displayData("deleteCommentTable",commentsData, "comment");
  } catch (error) {
    console.error('Error during GET request:', error.message);
  }
}




function displayData(tableId, data, type) {
  const table = document.getElementById(tableId);
  table.innerHTML = '';

  if (data.length === 0) {
    console.error('No data to display');
    return;
  }

  const keys = Object.keys(data[0]);

  // Check if "Close" button exists
  let closeButton = document.getElementById('closeButton');

  if (!closeButton) {
    // Add "Close" button only if it doesn't exist
    closeButton = document.createElement('button');
    closeButton.textContent = 'Close';
    closeButton.id = 'closeButton';
    closeButton.addEventListener('click', () => {
      table.innerHTML = ''; // Clear the table when the "Close" button is clicked
    });
    table.parentElement.insertBefore(closeButton, table);
  }

  // Header row
  const headerRow = table.insertRow(0);
  keys.forEach((key, index) => {
    const headerCell = headerRow.insertCell(index);
    headerCell.textContent = key;
  });

  // Add "Action" header
  const actionHeaderCell = headerRow.insertCell(keys.length);
  actionHeaderCell.textContent = 'Action';

  // Data rows
  data.forEach((rowData, rowIndex) => {
    const dataRow = table.insertRow(rowIndex + 1);
    keys.forEach((key, index) => {
      const dataCell = dataRow.insertCell(index);
      dataCell.textContent = rowData[key];
    });

    // Add "Delete" button in the last column
    const deleteButtonCell = dataRow.insertCell(keys.length);
    const deleteButton = document.createElement('button');
    const updButton = document.createElement('button');
    deleteButton.type = 'submit';

    switch (type) {
      case 'account':
        deleteButton.textContent = 'Delete Account';
        deleteButton.id = 'deleteAccount';
        deleteButton.addEventListener('click', () => deleteAccount(rowData.id)); 
        break;
      case 'accountUpd':
          updButton.textContent = 'Change Permissions';
          updButton.id = 'updAccount';
          updButton.addEventListener('click', function (event) {
            event.preventDefault();
            openAccessLevelModal(event, rowData);
        });
         
          break;
      case 'post':
        deleteButton.textContent = 'Delete Post';
        deleteButton.id = 'deletePost';
        deleteButton.addEventListener('click', () => deletePost(rowData.id));
        break;
      case 'comment':
        deleteButton.textContent = 'Delete Comment';
        deleteButton.id = 'deleteComment';
        deleteButton.addEventListener('click', () => deleteComment(rowData.id));
        break;
    }

    if(type === "accountUpd"){
      deleteButtonCell.appendChild(updButton);
    }else{
      deleteButtonCell.appendChild(deleteButton);
    }
    
  });

  table.classList.add('styled-table');
}



async function openAccessLevelModal(event, rowData) {
  event.preventDefault();
  const updAccountButton = document.getElementById('updAccount');
  const popupElement = document.getElementById('popup');
  const backdropElement = document.getElementById('backdrop');

  // Remove existing event listeners before adding new ones
  updAccountButton.removeEventListener('click', openAccessLevelModal);
  updAccountButton.removeEventListener('click', () => {
    // Populate modal with rowData
    document.getElementById('id').value = rowData.id;
    document.getElementById('fname').value = rowData.fname;
    document.getElementById('lname').value = rowData.lname;
    document.getElementById('fullname').value = rowData.fullname;
    document.getElementById('email').value = rowData.email;
    document.getElementById('pwd').value = rowData.pwd;
    document.getElementById('pnum').value = rowData.pnum;
    document.getElementById('username').value = rowData.username;
    document.getElementById('newAccessLevel').value = rowData.accesslevel;

    popupElement.classList.add('show');
    backdropElement.classList.add('show');
  });

  // Add new event listener
  updAccountButton.addEventListener('click', () => {
    // Populate modal with rowData
    document.getElementById('id').value = rowData.id;
    document.getElementById('fname').value = rowData.fname;
    document.getElementById('lname').value = rowData.lname;
    document.getElementById('fullname').value = rowData.fullname;
    document.getElementById('email').value = rowData.email;
    document.getElementById('pwd').value = rowData.pwd;
    document.getElementById('pnum').value = rowData.pnum;
    document.getElementById('username').value = rowData.username;
    document.getElementById('newAccessLevel').value = rowData.accesslevel;

    popupElement.classList.add('show');
    backdropElement.classList.add('show');
  });

  document.addEventListener('click', (event) => {
    if (!popupElement.contains(event.target) && event.target !== updAccountButton) {
      closePopup();
    }
  });

  document.addEventListener('keydown', (event) => {
    if (event.key === 'Escape') {
      closePopup();
    }
  });

  function closePopup() {
    popupElement.classList.remove('show');
    backdropElement.classList.remove('show');
  }

  document.getElementById('promoteButton').addEventListener('click', (event) => {
    event.preventDefault();
    promoteAccount(event, rowData);
  });
}


async function promoteAccount(event, rowData){
  event.preventDefault();
  const formData = new FormData(document.getElementById('updateAccessLevelModalForm'));
  const url = 'http://localhost:8080/account/update'

  const old_body = {
    id:		   rowData.id,
    fname:      rowData.fname,
    lname:      rowData.lname,
    fullname:   rowData.fullname,
    email:      rowData.email,
    pwd:        rowData.pwd,
    pnum:      rowData.pnum,
    username:   rowData.username,
    accesslevel: rowData.accesslevel,
  }
  const new_body ={
    id:		   rowData.id,
    fname:      rowData.fname,
    lname:      rowData.lname,
    fullname:   rowData.fullname,
    email:      rowData.email,
    pwd:        rowData.pwd,
    pnum:      rowData.pnum,
    username:   rowData.username,
    accesslevel: formData.get('newAccessLevel'),
   }

   try{
    const requestBody = {
        old: old_body,
        new: new_body,
    };
    console.table(requestBody)
    

    const response = await fetch(url, {
        method: 'PATCH',
        body: JSON.stringify(requestBody),
        headers: {
            'Content-Type': 'application/json',
        },
    });
    if(!response.ok){
        throw new Error(`HTTP error! Status: ${response.headers}`);
    }
    
    alert("Access Updated!")
    
}catch (error){
    console.error(`Error during patch method:`, error.message)
}
}