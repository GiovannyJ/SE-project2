let initialData; 

async function loadAccountData(){
    const url = 'http://localhost:8080/account';
    const response = await fetch(url);
    const accountData = await response.json();

    const tableBody = document.getElementById('accountTableBody');
    tableBody.innerHTML = '';

    accountData.forEach(account => {
        const row = tableBody.insertRow(-1);
        const keys = Object.keys(account);

        keys.forEach(key => {
            // Save initial data when rendering rows
            if (!initialData) {
                initialData = { ...account };
            }

            const cell = row.insertCell(-1);
            cell.textContent = account[key];
        });

        // Add Update button with onclick event
        const updateCell = row.insertCell(-1);
        const updateButton = document.createElement('button');
        updateButton.textContent = 'Update';
        updateButton.className = 'btn btn-warning';
        updateButton.onclick = function () {
            openUpdateModal(account);
        };
        updateCell.appendChild(updateButton);
    });
}
loadAccountData();

function openUpdateModal(account) {
    // Populate the modal form fields with current data
    const updateForm = document.getElementById('updateForm');
    updateForm.innerHTML = ''; // Clear previous form content

    Object.keys(account).forEach(key => {
        const label = document.createElement('label');
        label.textContent = key.charAt(0).toUpperCase() + key.slice(1);

        const input = document.createElement('input');
        input.type = 'text';
        input.value = account[key];
        input.name = key;

        // Make the "Id" column readonly
        if (key.toLowerCase() === 'id') {
            input.readOnly = true;
            input.disabled = true;
        }

        updateForm.appendChild(label);
        updateForm.appendChild(input);
        updateForm.appendChild(document.createElement('br'));
    });

    $('#updateModal').modal('show');
}


async function submitUpdate() {
    const updateForm = document.getElementById('updateForm');
    const formData = new FormData(updateForm);

    // Clone initialData using JSON
    const oldData = JSON.parse(JSON.stringify(initialData));
    const newData = {};

    formData.forEach((value, key) => {
        // Convert 'pnum' and 'id' to integers
        oldData[key] = (key === 'pnum' || key === 'id') ? parseInt(value, 10) : value;
        newData[key] = (key === 'pnum' || key === 'id') ? parseInt(value, 10) : value;
    });

    const requestBody = { old: oldData, new: newData };

    const url = 'http://localhost:8080/account/update';
    const response = await fetch(url, {
        method: 'PATCH',
        body: JSON.stringify(requestBody),
        headers: {
            'Content-Type': 'application/json',
        },
    });

    if (response.ok) {
        // Handle success, e.g., close modal, refresh data
        $('#updateModal').modal('hide');
        loadAccountData();
    } else {
        console.error('Error during PATCH request:', response.status);
    }
}