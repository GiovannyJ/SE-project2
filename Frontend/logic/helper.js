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

    if (!userExists()) {
        document.location.href = 'login.html';
        return;
    }
  
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
        { label: "Admin Panel", href: "admin_panel.html" },
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
            <a href="searchresults.html"><img src="logo2.png" class="logo"></a>
            <ul class="rest">
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
    <a href="searchresults.html"><img src="logo2.png" class="logo"></a>
    <ul class="rest">
        <li>
            <div id="search-box">
                <select id="search-options">
                    <option value="id">ID</option>
                    <option value="title">Title</option>
                    <option value="genre">Genre</option>
                    <option value="authorID">Author ID</option>
                    <option value="numUp">Num Up</option>
                    <option value="numDown">Num Down</option>
                    <option value="postedDate">Posted Date</option>
                    <option value="order">Order</option>
                </select>
                <input type="text" id="search-input" placeholder="Search...">
                <button id="search-button" onclick="searchPosts()">Search</button>
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

function searchPosts() {
    const searchOptions = document.getElementById('search-options');
    const searchInput = document.getElementById('search-input').value;
    const selectedOption = searchOptions.options[searchOptions.selectedIndex].value;

    // Build the query string
    const queryString = encodeURIComponent(selectedOption) + '=' + encodeURIComponent(searchInput);

    // Redirect to searchresults.html with the query string
    window.location.href = 'searchresults.html?' + queryString;
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
    const user = JSON.parse(localStorage.getItem('user'));
    // Treat non-existing users as guests
    if (user === null) {
        return "guest";
    }
    return user.accesslevel;
}

function redirectMenu(){
    alert('Your account does not have access to this page!')
    window.location.href = 'menu.html';
}


function renderAdminPanel() {
    const pannelList = document.getElementById('adminPanelList');
    const accesslevel = getUserAccessLevel();

    if (accesslevel === "verified user") {
        const deleteAccountDiv = document.getElementById('deleteAccount');
        const changeAccessLvlDiv = document.getElementById('changeAccess')
        if (deleteAccountDiv && changeAccessLvlDiv) {
            pannelList.removeChild(deleteAccountDiv);
            pannelList.removeChild(changeAccessLvlDiv);
        }
    }
}