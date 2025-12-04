const params = new URLSearchParams(window.location.search)
const token = params.get('token')

document.addEventListener('DOMContentLoaded', () => {
  const confirmForm = document.getElementById('confirmForm')
  const inputCode = document.getElementById('inputCode')
  const errorMessage = document.getElementById('errorMessage')


  confirmForm.addEventListener('submit', async function (e) {
    e.preventDefault()

    const errorMessage = document.getElementById('errorMessage')
    const code = inputCode.value.trim()

    // Check code length
    const error = notValidCode(code)
    if (error) {
      errorMessage.textContent = error
      errorMessage.style.display = 'block'
      return
    }
    
    const response = await fetch(window.APP_CONFIG.url_confirm_code, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ code, token })
      })

      if (response.ok) {
        window.location.href = window.APP_CONFIG.url_app
        return
      }
      
      try {
        const dataError = await response.json()
        errorMessage.textContent = dataError.message
        errorMessage.style.display = 'block'
      }
      catch {
        console.error("Error parsing Json error")
        errorMessage.textContent = 'Error en el sistema. Intenta nuevamente mas tarde.'
      }
      errorMessage.style.display = 'block'

  })
})


function notValidCode(code) {
  if (code.length !== 6) {
    return 'El codigo debe tener 6 numeros.'
  }
  if (!/^\d{6}$/.test(code)) {
    return 'El codigo debe contener solo numeros.'
  }
  return ''
}