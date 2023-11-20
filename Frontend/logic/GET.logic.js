/**
 * ================================================================== 
* GET METHODS
* these methods are used to retrieve data from certain endpoints
* the data is binded in the response.json()
* using query string element can return specific elements 
* ==================================================================
*/

/**
 * gets all the accounts from the database
*/
async function getAccount() {
    var queryString = document.getElementById('queryString').value;
    var url = 'http://localhost:8080/account';
    
    if (queryString){
      url = 'http://localhost:8080/account?' + queryString;
    }
  
    try {
      const response = await fetch(url);
  
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
  
      const accountData = await response.json();
      displayData(accountData);
    } catch (error) {
      console.error('Error during GET request:', error.message);
    }
  }
  
  /**
   * gets all the posts in full context
  */
  async function getFullContextPosts() {
    var queryString = document.getElementById('queryString').value;
    var url = 'http://localhost:8080/posts/fullcontext?';
    
    if (queryString){
      url = url + queryString;
    }
  
    try {
      const response = await fetch(url);
  
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
  
      const postsData = await response.json();
      displayData(postsData);
    } catch (error) {
      console.error('Error during GET request:', error.message);
    }
  }
  
  /**
   * gets all the posts in the database 
  */
  async function getPosts() {
    var queryString = document.getElementById('queryString').value;
    var url = 'http://localhost:8080/posts?';
  
    if (queryString){
      url = url + queryString;
    }
  
    try {
      const response = await fetch(url);
  
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
  
      const postsData = await response.json();
      displayData(postsData);
    } catch (error) {
      console.error('Error during GET request:', error.message);
    }
  }
  
  /**
   * gets all the comments relating to a post using postID
  */
  async function getComments() {
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
  
  /**
   * gets comments in full context using postID
   */
  async function getCommentsFullContext() {
    const postId = document.getElementById('postId').value;
    var queryString = document.getElementById('queryString').value;
    var url = `http://localhost:8080/posts/${postId}/commentsfullcontext?`;
  
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
  
  /**
   * displays image from database using imagename
   */
  function displayImage() {
    var filename = document.getElementById('filename').value;
    var imageContainer = document.getElementById('imageContainer');
    imageContainer.innerHTML = ''; // Clear previous image
  
    var imageElement = document.createElement('img');
    imageElement.src = 'http://localhost:8080/uploads/' + filename;
    imageElement.alt = 'Image not found'; // Alt text if the image is not found
    imageContainer.appendChild(imageElement);
  }
  
  /**
   * helper method to display data of GET method
   */
  function displayData(data) {
    const tableHeader = document.getElementById('table-header');
    const tableBody = document.getElementById('table-body');
  
    // Clear previous data
    tableHeader.innerHTML = '';
    tableBody.innerHTML = '';
  
    // Display table headers
    const dataKeys = Object.keys(data[0]);
    dataKeys.forEach(key => {
      tableHeader.innerHTML += `<th>${key}</th>`;
    });
  
    // Display data
    data.forEach(entry => {
      const row = document.createElement('tr');
      dataKeys.forEach(key => {
        const cell = document.createElement('td');
  
        // Check if the property is nested
        if (typeof entry[key] === 'object') {
          // Access nested properties
          const nestedKeys = Object.keys(entry[key]);
          nestedKeys.forEach(nestedKey => {
            cell.innerHTML += `<p>${nestedKey}: ${entry[key][nestedKey]}</p>`;
          });
        } else {
          cell.textContent = entry[key];
        }
  
        row.appendChild(cell);
      });
      tableBody.appendChild(row);
    });
  }