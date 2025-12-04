document.getElementById('registerForm').addEventListener('submit', async function (e) {
  const formContainer = document.getElementById('formContainer')
  const successfulContainer = document.getElementById('successfulContainer')
  const errorMessage = document.getElementsByClassName('errorMessage')[0]
  e.preventDefault()

  const params = new URLSearchParams(window.location.search)
  const token = params.get('token')

  const email = document.getElementById('email').value

  const response = await fetch(window.APP_CONFIG.url_register, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ email, token })

    })
    if (response.ok) {
      const data = await response.json()
      console.log("Registered:", data)
      errorMessage.style.display = 'none'
      formContainer.style.display = 'none'
      successfulContainer.style.display = 'block'
      return
    }
    else {
      try {
        const dataError = await response.json()
        errorMessage.textContent = dataError.message
        errorMessage.style.display = 'block'
        
      }
      catch {
        console.error("Error parsing Json error")
        errorMessage.textContent = 'Error en el sistema. Intenta nuevamente mas tarde.'
        errorMessage.style.display = 'block'
      }
    }
})


