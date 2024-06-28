const url = "http://0.0.0.0:8080/login"

document.getElementById('loginForm').addEventListener('submit', async function(event) {
  event.preventDefault();
  const resposta = await fetch(url)

  console.log(resposta)
  //
  //let data = { username: document.getElementById('username'), password: document.getElementById('password') };

  //fetch('', {
  //  // method: 'POST',
  //  // headers: {
  //  //   'Content-Type': 'application/json',
  //  // },
  //  // boby: JSON.stringify(data),
  //})
  //  //.then(response => response.json())
  //  .then(data => {
  //    console.log('resposta do backend', data);
  //  })
  //  .catch(error => {
  //    console.error('erro ao enviar os dados', error);
  //  })
});
