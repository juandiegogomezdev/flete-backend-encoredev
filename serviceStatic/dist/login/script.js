document.getElementById('loginForm').addEventListener('submit', async function (e) {
  e.preventDefault()

  const email = document.getElementById('email').value
  const password = document.getElementById('password').value
  const errorMessage = document.querySelector('.errorMessage')
  try {
    const response = await fetch(window.APP_CONFIG.url_login, {
      method: 'POST',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ email, password })

    })


    if (response.ok) {
      const data = await response.json()
      window.location.href = window.APP_CONFIG.page_url_login_confirm+'?token=' + data.token

    }
    else {
      try {
        const errorData = await response.json()
        errorMessage.textContent = errorData.message
        errorMessage.style.display = 'block'

      }
      catch {
        console.error("Error parsing Json error")
        errorMessage.textContent = 'Error en el sistema. Intent nuevamente mas tarde.'

      }
      errorMessage.style.display = 'block'
    }
    
  } catch (error) {
    console.error(error)
    errorMessage.style.display = 'block'
    errorMessage.textContent = 'Error en el sistema . Intent nuevamente mas tarde.'
  }
})
