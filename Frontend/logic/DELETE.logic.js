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
export async function deleteAccount() {
  const userConfirmed = window.confirm("Are you sure you would like to delete this account?");
    if (!userConfirmed) {
        alert("Account was not deleted");
        return;
    } 
    const accountId = document.getElementById('accountSearchInput').value;
    const url = `http://localhost:8080/account/delete/${accountId}`;
  
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
  export async function deletePost() {
    const userConfirmed = window.confirm("Are you sure you would like to delete this Post?");
    if (!userConfirmed) {
        alert("Post was not deleted");
        return;
    } 

    const postId = document.getElementById('postSearchInput').value;
    const url = `http://localhost:8080/posts/delete/${postId}`;
  
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
  export async function deleteComment() {
    const commentId = document.getElementById('commentId').value;
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

export async function displayAccount(event, type) {
    event.preventDefault();
    var url = 'http://localhost:8080/account';
    const param = document.getElementById('accountSearchInput').value;

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
            displayData('promoteAccountTable', accountData, "account");
            break;
        }

    } catch (error) {
        console.error('Error during GET request:', error.message);
    }
}

  
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
    
    displayData('deletePostTable', postsData);
    
  } catch (error) {
    console.error('Error during GET request:', error.message);
  }
}

export async function getComments() {
  const postId = document.getElementById('postId').value;
  var queryString = document.getElementById('queryString').value;
  var url = `http://localhost:8080/posts/${postId}/comments?`;

  if (queryString){
    url = url + queryString;
  }

  try {
    const response = await fetch(url);

    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }

    const commentsData = await response.json();
    displayData(commentsData);
  } catch (error) {
    console.error('Error during GET request:', error.message);
  }
}











function displayData(tableId, data) {
    const table = document.getElementById(tableId);
    table.innerHTML = '';

    if (data.length === 0) {
        console.error('No account data to display');
        return;
    }

    const keys = Object.keys(data[0]);

    // Header row
    const headerRow = table.insertRow(0);
    keys.forEach((key, index) => {
        const headerCell = headerRow.insertCell(index);
        headerCell.textContent = key;
    });

    // Data row
    const dataRow = table.insertRow(1);
    keys.forEach((key, index) => {
        const dataCell = dataRow.insertCell(index);
        dataCell.textContent = data[0][key];
    });

    table.classList.add('styled-table')
}