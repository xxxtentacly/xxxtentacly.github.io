const registrationForm = document.getElementById('registration-form');
const loginForm = document.getElementById('login-form');
const message = document.getElementById('message');

registrationForm.addEventListener('submit', async (event) => {
	event.preventDefault();
	const formData = new FormData(registrationForm);
	const username = formData.get('username');
	const email = formData.get('email');
	const password = formData.get('password');
	try {
		const response = await fetch('/register', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ username, email, password })
		});
		const data = await response.json();
		message.innerHTML = data.message;
	} catch (error) {
		console.error(error);
		message.innerHTML = 'Ошибка при регистрации';
	}
});

loginForm.addEventListener('submit', async (event) => {
	event.preventDefault();
	const formData = new FormData(loginForm);
	const username = formData.get('username');
	const password = formData.get('password');
	try {
		const response = await fetch('/login', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ username, password })
		});
		const data = await response.json();
		message.innerHTML = data.message;
	} catch (error) {
		console.error(error);
		message.innerHTML = 'Ошибка при входе';
	}
});
