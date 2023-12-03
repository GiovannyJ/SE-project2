/**
 * HELPER METHODS FOR THE ENTIRE CODEBASE
 */



  /**
   * clears local storage and returns to login page
   */
function signOut() {
    localStorage.clear();
    window.location = 'login.html'
}

  /**
   * renders the menu depending on user access level
   * 
   */
function renderMenu() {
    const user = JSON.parse(localStorage.getItem('user'));
    const menuContainer = document.getElementById('menu');
  
    // Clear existing menu items
    menuContainer.innerHTML = '<ul></ul>';
  
    switch (user.accesslevel) {
      case "admin":
        renderAdminMenu(menuContainer);
        break;
      case "guest":
        renderGuestMenu(menuContainer);
        break;
      case "user":
        renderUserMenu(menuContainer);
        break;
      case "verified user":
        renderVerifiedUserMenu(menuContainer);
        break;
      default:
        console.error("Unknown access level");
        break;
    }
  }

  /**
 * all admin can see
 */
function renderAdminMenu(container) {
    const menuItems = [
        { label: "Home", href: "searchresults.html" },
        { label: "Ask a Question", href: "ask.html" },
        { label: "Account Info", href: "updateAccount.html" },
        { label: "Admin Panel", href: "admin_panel.html" },
    ];

    renderMenuItems(container, menuItems);
}

/**
 * renders what guest can see
 */
function renderGuestMenu(container) {
    const menuItems = [
        { label: "Home", href: "searchresults.html" },
    ];

    renderMenuItems(container, menuItems);
}
  
/**
 * renders what user can see
 */
function renderUserMenu(container) {
    const menuItems = [
        { label: "Home", href: "searchresults.html" },
        { label: "Ask a Question", href: "ask.html" },
        { label: "Account Info", href: "updateAccount.html" },
        ];

    renderMenuItems(container, menuItems);
}
 
/**
 * renders what verified user can see
 */
function renderVerifiedUserMenu(container) {
    const menuItems = [
        { label: "Home", href: "searchresults.html" },
        { label: "Ask a Question", href: "ask.html" },
        { label: "Account Info", href: "updateAccount.html" },
        //{ label: "Admin Panel", href: "admin_panel.html" },
        ];

    renderMenuItems(container, menuItems);
}
 
/**
 * renders generic menu
 */
function renderMenuItems(container, items) {
    const ul = container.querySelector('ul');

    items.forEach(item => {
        const li = document.createElement('li');
        const div = document.createElement('div');
        const a = document.createElement('a');
        const h1 = document.createElement('h1');

        a.href = item.href;
        h1.textContent = item.label;

        div.classList.add('menuItems');
        div.appendChild(a);
        a.appendChild(h1);
        li.appendChild(div);
        ul.appendChild(li);
    });
}

/**
 * renders header with sign in button
 * to be used on pages that the user can see without singing in
 */
function renderHeader_signIn(){
    const header = document.createElement('header');
    header.innerHTML  = `
        <nav>
            <a href="#"><img src="logo2.png" class="logo"></a>
            <ul class="rest">
                <li>
                    <div id="search-box">
                        <input type="text" id="search-input" placeholder="Search...">
                        <button id="search-button" onclick="searchDatabase()">Search</button>
                        <!--need to create searchDatabase() to search database for query-->
                    </div>
                </li>
                <li><a href="about.html" class="nav-button1">About Us</a></li>
                <li><a href="support.html" class="nav-button1">Support</a></li>
                <li><a href="login.html" class="nav-button2">Sign In</a></li>
                <li><a href="createacct.html" class="nav-button2">Create Account</a></li>
            </ul>
        </nav>
        `
    const parentElement = document.querySelector('main');

    addChildToFront(parentElement, header);
}

/**
 * renders page with sign out button
 * to be used on pages where user has to be logged in to see
 */
function renderHeader_signOut(){
    const header = document.createElement('header');
    header.innerHTML = `
        <nav>
            <a href="#"><img src="logo2.png" class="logo"></a>
            <ul class="rest">
                <li>
                    <div id="search-box">
                        <input type="text" id="search-input" placeholder="Search...">
                        <button id="search-button" onclick="getPosts()">Search</button>
                    </div>
                </li>
                <li><a href="about.html" class="nav-button1">About Us</a></li>
                <li><a href="support.html" class="nav-button1">Support</a></li>
                <li><a href="menu.html" class="nav-button2">Dashboard</a></li>
                <li><button onclick="signOut()" class="nav-button2">Sign Out</button></li>
            </ul>
        </nav>
    `;

    const parentElement = document.querySelector('main');

    addChildToFront(parentElement, header);
}

//helper method to make sure that header renders in correct spot before main
function addChildToFront(parent, newChild) {
    parent.insertBefore(newChild, parent.firstChild);
}


/**
 * check if user exists in local storage
 * @returns true if user exists and false if not
 */
function userExists() {
    const user = JSON.parse(localStorage.getItem('user'));
    return user !== null;
}

function getUserAccessLevel(){
    const user = JSON.parse(localStorage.getItem('user'))
    return user.accesslevel;
}

function redirectMenu(){
    alert('Your account does not have access to this page!')
    window.location.href = 'menu.html';
}



function renderAdminPanel(){
    const pannelList = document.getElementById('adminPanelList')

    accesslevel = getUserAccessLevel()

    switch(accesslevel){
        case "admin":
            pannelList.innerHTML = `
            <div class="admin-control">
                <form id="deleteAccForm">
                    <h2><u>Delete Account</u></h2>
                    
    
                    <table id="deleteAccountTable">
                        <thead>
                           
                        </thead>
                        <tbody>
                           
                        </tbody>
                    </table>
                    <br>
                    <input type="text" id="accountSearchInput" placeholder="Search Account">
                    <br>
                    <br>
                    <button id="searchAccount">
                        Search Account
                    </button>
                </form>
            </div>
            <br>
        
            <div class="admin-control">
                <form id="deletePostForm">
                    <h2><u>Delete Post</u></h2>
                    
                    <table id="deletePostTable">
                        <thead>
                           
                        </thead>
                        <tbody>
                            
                        </tbody>
                    </table>
                    <input type="text" id="postSearchInput" placeholder="Search Post">
                    <br>
                    <br>
                    <button id="searchPost">
                        Search Post
                    </button>
                </form>
            </div>
            <br>
        
            <div class="admin-control">
                <form id="deleteCommentForm">
                    <h2><u>Delete Comment</u></h2>
                    
                    <table id="deleteCommentTable">
                        <thead>
                           
                        </thead>
                        <tbody>
                            
                        </tbody>
                    </table>
                    <br>
                    
                    <input type="text" id="postID" placeholder="Post ID" required><strong style="color:red;">*<sup style="font-size: x-small;">Required</sup></strong>
                    <br>
                    <br>
                    <input type="text" id="commentSearchInput" placeholder="Search Comment">
                    <br>
                    <br>
                    <button id="searchComment">
                        Search Comment
                    </button>
                </form>
            </div>
            <br>

            <div class="admin-control">
                <form id="changeAccessLvlForm">
                    <h2><u>Change Access Level</u></h2>
                    
                    <table id="promoteAccTable">
                        <thead>
                            
                        </thead>
                        <tbody>
                            
                        </tbody>
                    </table>
                    <input type="text" id="promoteAcc" placeholder="Search Account">
                    <br>
                    <br>
                    <button id="searchAcc">
                        Search Account
                    </button>
                </form>
            </div>
        </ul>

        <div id="popup">
            <form id="updateAccessLevelModalForm">
                
                <div>
                    id:<input type="text" id="id" readonly>	
                    first name:<input type="text" id="fname" readonly>
                    last name: <input type="text" id="lname" readonly>
                    fullname: <input type="text" id="fullname" readonly>
                    email: <input type="text" id="email" readonly>
                    password: <input type="text" id="pwd" readonly>
                    phone number:<input type="text" id="pnum" readonly>
                    username:<input type="text" id="username" readonly>
                    Access Level:
                    <select id="newAccessLevel" name="newAccessLevel" placeholder="Select Privilege">
                        <option value="0">Select Priv</option>
                        <option value="user">User</option>
                        <option value="verified user">Verified User</option>
                        <option value="admin">Admin</option>
                    </select>
                </div>
                <button type="submit" id="promoteButton">
                    Change Access Level
                </button>
            </form>
        </div>
    `
    break;
    case "verified user":
        pannelList.innerHTML = `
        <br>
    
        <div class="admin-control">
            <form id="deletePostForm">
                <h2><u>Delete Post</u></h2>
                
                <table id="deletePostTable">
                    <thead>
                       
                    </thead>
                    <tbody>
                        
                    </tbody>
                </table>
                <input type="text" id="postSearchInput" placeholder="Search Post">
                <br>
                <br>
                <button id="searchPost">
                    Search Post
                </button>
            </form>
        </div>
        <br>
    
        <div class="admin-control">
            <form id="deleteCommentForm">
                <h2><u>Delete Comment</u></h2>
                
                <table id="deleteCommentTable">
                    <thead>
                       
                    </thead>
                    <tbody>
                        
                    </tbody>
                </table>
                <br>
                
                *<input type="text" id="postID" placeholder="Post ID" required>
                <br>
                <br>
                <input type="text" id="commentSearchInput" placeholder="Search Comment">
                <br>
                <br>
                <button id="searchComment">
                    Search Comment
                </button>
            </form>
        </div>
        <br>
    `
    break;
    }



    
}


