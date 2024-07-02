const url = "http://localhost:8080/post"

document.getElementById('login').addEventListener('submit', async function () {
  const username = document.getElementById('username').value
  const password = document.getElementById('password').value

  var myHeaders = new Headers()
  myHeaders.append("Content-Type", "application/json")

  var myInit = {
    method: "POST",
    Headers: myHeaders,
    mode: "cors",
    cache: "default",
    boby: JSON.stringify({ "username": username, "password": password })
  }

  fetch(url, myInit)
    .then(function (response) {
      if (response.ok) {
        console.log("tudo OK!")
        console.log(response.statusText)
      } else {
        console.log("API respondeu not ok")
        console.log(response.statusText)
      }
    }).catch(function (error) {
      console.log("algo deu errado: " + error.message)
    })
})
