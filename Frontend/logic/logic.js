async function getAccount() {
    const url = 'http://localhost:8080/account';

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

  async function getFullContextPosts() {
    const url = 'http://localhost:8080/posts/fullcontext';

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

  async function getPosts() {
    const url = 'http://localhost:8080/posts';

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
        cell.textContent = entry[key];
        row.appendChild(cell);
      });
      tableBody.appendChild(row);
    });
  }

  async function deleteAccount() {
    const accountId = document.getElementById('accountId').value;
    const url = `http://localhost:8080/account/delete/${accountId}`;

    try {
      const response = await fetch(url, {
        method: 'DELETE',
      });

      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      alert('Account deleted successfully!');
    } catch (error) {
      console.error('Error during DELETE request:', error.message);
    }
  }

  async function deletePost() {
    const postId = document.getElementById('postId').value;
    const url = `http://localhost:8080/posts/delete/${postId}`;

    try {
      const response = await fetch(url, {
        method: 'DELETE',
      });

      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      alert('Post deleted successfully!');
    } catch (error) {
      console.error('Error during DELETE request:', error.message);
    }
  }

  async function createAccount() {
    const formData = new FormData(document.getElementById('accountForm'));
    const url = 'http://localhost:8080/account/new';

    try {
      const requestBody = {
        Fname: formData.get('Fname'),
        Lname: formData.get('Lname'),
        Fullname: formData.get('Fullname'),
        Email: formData.get('Email'),
        pwd: formData.get('pwd'),
        Pnum: formData.get('Pnum'),
        Username: formData.get('Username'),
      };

      const response = await fetch(url, {
        method: 'POST',
        body: JSON.stringify(requestBody),
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      const result = await response.json();
      displayResult('accountTable', result);
    } catch (error) {
      console.error('Error during POST request:', error.message);
    }
  }

  async function createPost() {
    const formData = new FormData(document.getElementById('postForm'));
    const url = 'http://localhost:8080/posts/create';

    try {
      const requestBody = {
        id: formData.get('id'),
        title: formData.get('title'),
        descr: formData.get('descr'),
        genre: formData.get('genre'),
        authorId: formData.get('authorId'),
        picId: formData.get('picId'),
      };

      const response = await fetch(url, {
        method: 'POST',
        body: JSON.stringify(requestBody),
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      const result = await response.json();
      displayResult('postTable', result);
    } catch (error) {
      console.error('Error during POST request:', error.message);
    }
  }

  async function login() {
    const formData = new FormData(document.getElementById('loginForm'));
    const url = 'http://localhost:8080/login';

    try {
      const requestBody = {
        username: formData.get('username'),
        password: formData.get('password'),
        email: formData.get('email'),
      };

      const response = await fetch(url, {
        method: 'POST',
        body: JSON.stringify(requestBody),
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      const result = await response.json();
      displayResult('loginTable', result);
    } catch (error) {
      console.error('Error during POST request:', error.message);
    }
  }

  function displayResult(tableId, result) {
    const table = document.getElementById(tableId);
    table.innerHTML = '';

    const keys = Object.keys(result);
    const headerRow = table.insertRow(0);
    keys.forEach(key => {
      const headerCell = headerRow.insertCell(-1);
      headerCell.textContent = key;
    });

    const dataRow = table.insertRow(1);
    keys.forEach(key => {
      const dataCell = dataRow.insertCell(-1);
      dataCell.textContent = result[key];
    });
  }
  