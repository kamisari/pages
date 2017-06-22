function init() {
  document.getElementById('hi').onclick = function() { alert('hi') }

  let togglediv = document.querySelector('div')
  togglediv.onclick = function() {
    let attr = togglediv.getAttribute('class')
    if (attr === 'back1') {
      togglediv.setAttribute('class', 'back2')
    } else {
      togglediv.setAttribute('class', 'back1')
    }
  }
}

