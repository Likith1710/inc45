document.addEventListener('DOMContentLoaded', function() {
    const form = document.getElementById('nameForm');

    form.addEventListener('submit', function(event) {
        event.preventDefault();
        const nameInput = document.getElementById('name');
        const name = nameInput.value.trim();

        if (name !== '') {
            // Perform calculations and display results
            const result = 'Addition: 10, Multiplication: 20'; // Replace with actual calculation result
            alert(`Hello, ${name}!\n${result}`);

            // Clear input field
            nameInput.value = '';
        } else {
            alert('Please enter your name.');
        }
    });
});

