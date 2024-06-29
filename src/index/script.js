const url = "http://0.0.0.0:8080/login"

document.getElementById('loginForm').addEventListener('submit', async function(event) {
  event.preventDefault();
  const username: document.getElementById('username').value;

  const password: document.getElementById('password').value;

  console.log(username)
  console.log(password)

  try {
    const response = await fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      boby: JSON.stringify({ username, password }),
    });

    if (!response.ok) {
      throw new Error('Erro na requisição: ' + response.statusText);
    }
    const responseData = await response.json();
    console.log('Resposta do backend:', responseData);

  } catch (error) {
    console.error('Erro ao enviar os dados:', error);
  }
});
