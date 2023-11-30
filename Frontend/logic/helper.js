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
        { label: "Moderator Panel", href: "moderator_panel.hml"}
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
        { label: "Moderator Panel", href: "moderator_panel.hml"}
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
                <li><button onclick="signOut()" class="nav-button2">Sign Out</button></li>
                <li><a href="createacct.html" class="nav-button2">Create Account</a></li>
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
 * @returns true if user exists and fasle if not
 */
function userExists() {
    const user = JSON.parse(localStorage.getItem('user'));
    return user !== null;
}
