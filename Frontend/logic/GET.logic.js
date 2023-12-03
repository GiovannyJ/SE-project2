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
export async function getAccount(param) {
  var url = 'http://localhost:8080/account';

  if (param) {
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

/**
 * gets all the posts in full context
*/
export async function getFullContextPosts(postID) {
  // var queryString = document.getElementById('queryString').value;
  var url = `http://localhost:8080/posts/fullcontext?postID=${postID}`;

  // if (queryString){
  //   url = url + queryString;
  // }

  try {
    const response = await fetch(url);

    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }

    const postsData = await response.json();
    displayPostFullContext(postsData);
  } catch (error) {
    console.error('Error during GET request:', error.message);
  }
}

function displayPostFullContext(postsData) {
  const post = postsData[0];

  document.getElementById('postID').value = post.id;
  //document.getElementById('author').textContent = "Author: " + post.authorInfo.username;
  document.getElementById('title').textContent = post.title;
  document.getElementById('genre').textContent = "Genre: " + post.genre;
  //document.getElementById('date').textContent = "posted: " + post.postedDate;
  document.getElementById('author').textContent = "Posted by "+post.authorInfo.username+" on "+post.postedDate;

  // Check if numUp is null, and set a default value of 0
  document.getElementById('numUp').value = post.numUp !== null ? post.numUp : 0;

  // Check if numDown is null, and set a default value of 0
  document.getElementById('numDown').value = post.numDown !== null ? post.numDown : 0;

  document.getElementById('descr').innerHTML = post.descr;

  const imgName = post.picInfo.imgname;
  if (imgName) {
    const imgElement = document.getElementById('img');

    const imgHtml = `<img src='../uploads/${imgName}' alt='img not found' height='15%'; width='auto';>`;

    imgElement.innerHTML = imgHtml;
  }
}

/**
 * gets all the posts in the database 
*/
export async function getPosts() {
  var url = 'http://localhost:8080/posts?';

  try {
    const response = await fetch(url);

    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }

    const postsData = await response.json();
    displayPost(postsData);
  } catch (error) {
    console.error('Error during GET request:', error.message);
  }
}

function displayPost(postsData) {
  const resultsContainer = document.getElementById("results-container");

  // Clear existing content
  resultsContainer.innerHTML = '';

  // Loop through the postsData array and create elements for each post
  postsData.forEach((post) => {
    const listItem = document.createElement("li");
    const postDiv = document.createElement("div");
    const postLinkButton = document.createElement("button");
    const profileButton = document.createElement("button");
    const genreHeading = document.createElement("h6");
    const titleHeading = document.createElement("h3");
    const authorHeading = document.createElement("h6");
    const dateHeading = document.createElement("h6");
    const likesHeading = document.createElement("h6");
    const descriptionParagraph = document.createElement("p");


    postLinkButton.textContent = "View Post";
    postLinkButton.id = post.id;
    postLinkButton.classList.add("nav-button2"); // Add the button-like class
    postLinkButton.onclick = function () {
      window.location.href = 'question.html?id=' + post.id;
    };

    titleHeading.textContent = post.title;
    genreHeading.textContent = "Genre: " + post.genre;
    authorHeading.textContent = "Posted By: " + post.authorId;
    profileButton.textContent = "PROFILE"
    profileButton.id = "profile" + post.authorId;
    descriptionParagraph.textContent = post.descr;
    dateHeading.textContent = "Date Posted: " + post.postedDate;
    likesHeading.textContent = "Upvotes: " + post.numUp;
 
    // Append elements to the postDiv
    //postDiv.appendChild(document.createElement("br"));
    postDiv.appendChild(titleHeading);
    postDiv.appendChild(genreHeading);
    //postDiv.appendChild(profileButton);
    postDiv.appendChild(descriptionParagraph);
    postDiv.appendChild(authorHeading);
    postDiv.appendChild(dateHeading);
    postDiv.appendChild(likesHeading);
    postDiv.appendChild(postLinkButton);

    // Set class for styling
    postDiv.classList.add("searchResults");

    // Append postDiv to the listItem, and listItem to the resultsContainer
    listItem.appendChild(postDiv);
    resultsContainer.appendChild(listItem);
  });
}



/**
 * gets all the comments relating to a post using postID
*/
async function getComments() {
  const postId = document.getElementById('postId').value;
  var queryString = document.getElementById('queryString').value;
  var url = `http://localhost:8080/posts/${postId}/comments?`;

  if (queryString) {
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
export async function getCommentsFullContext() {
  const postId = document.getElementById('postID').value;
  // const postId = 14;
  // var queryString = document.getElementById('queryString').value;
  var url = `http://localhost:8080/posts/${postId}/commentsfullcontext?`;

  // if (queryString){
  //   url = url + queryString;
  // }

  try {
    const response = await fetch(url);

    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }

    const commentsData = await response.json();
    displayComments(commentsData);
  } catch (error) {
    console.error('Error during GET request:', error.message);
  }
}

function displayComments(commentsData) {
  const commenterElement = document.getElementById('commenter');

  // Clear existing comments
  commenterElement.innerHTML = '';

  // Iterate through commentsData and append each comment to the 'commenter' element
  commentsData.forEach(comment => {
    const commentContainer = document.createElement('div');
    // commentContainer.classList.add('comment-container');
    commentContainer.id = `comment${comment.commentInfo.id}`; // Assuming each comment has a unique ID
    commentContainer.innerHTML = `
            <ul>
                <li><img src="pfp.png" id="pfp" alt="circle"></li>
                <li>
                    <div>
                        <p class="comment-container">${comment.commentInfo.content}</p>
                        <p>Posted by ${comment.commenterInfo.username} on ${comment.commentInfo.postedDate}</p>
                        <!--<p>Commented Date: ${comment.commentInfo.postedDate}</p>-->
                        <p>${comment.commentInfo.numUp} upvotes</p>
                        <p>${comment.commentInfo.numDown} downvotes</p>
                    </div>
                </li>
            </ul>
            <input type="hidden" id="commentId" value="${comment.commentInfo.id}">
        `;

    commenterElement.appendChild(commentContainer);
  });
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



