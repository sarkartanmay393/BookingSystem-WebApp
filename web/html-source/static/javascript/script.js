
// Script for form validation
(function () {
    'use strict'
    const forms = document.querySelectorAll('.needs-validation')
    Array.from(forms)
        .forEach(function (form) {
            form.addEventListener('submit', function (event) {
                if (!form.checkValidity()) {
                    event.preventDefault()
                    event.stopPropagation()
                }

                form.classList.add('was-validated')
            }, false)
        })
})()

// Script for date range picker
const elem = document.getElementById('reservation-dates');
const rangePicker = new DateRangePicker(elem, {
    // Put Options Here
    format: "dd-mm-yyyy"
});

// Script for alert on screen using Notie package
function notify(messageText, messageType) {
    notie.alert({
        type: messageType, // optional, default = 4, enum: [1, 2, 3, 4, 5, 'success', 'warning', 'error', 'info', 'neutral']
        text: messageText,
        position: "bottom"
        // stay: Boolean, // optional, default = false
        // time: Number, // optional, default = 3, minimum = 1,
    })
}

// Script for prompt alert using SweetAlert
function notifySweet(title, text, icon, confirmButtonText) {
    Swal.fire({
        title: title,
        text: text,
        icon: icon,
        confirmButtonText: confirmButtonText
    });
}

let buttonAttention = Prompt();

// JS code for a specific button on how to react of click.
document.getElementById("test-btn").addEventListener("click", function () {
    let html = `
      <form action="" method="post" class="needs-validation" novalidate>
          <div id="reservation-dates-popup" class="row g-2 mt-2">
            <div class="col">
              <div class="mb-3">
                <input type="text" class="form-control" id="start-date" name="start-date" aria-describedby="start-date-help" placeholder="Arrival Date" required>
                <div class="invalid-feedback">Enter Valid Arrival Date</div>
              </div>
            </div>
            <div class="col">
              <div class="mb-3">
                <input type="text" class="form-control" id="end-date" name="end-date" aria-describedby="end-date-help" placeholder="Departure Date" required>
                <div class="invalid-feedback">Enter Valid Departure Date</div>
              </div>
            </div>
          </div>
      </form>
      `
    // notify("This is my msg", "success")
    // notifySweet("error!", "do you test", "error", "okay");
    buttonAttention.formPopup({msg: html, title: "Choose dates"})
})



// JS function for Pop-up
function Prompt() {
    let toast = function (c) {
        const {
            msg = "",
            icon = "success",
            position = "top-end",
        } = c;
        const Toast = Swal.mixin({
            toast: true,
            title: msg,
            position: position,
            icon: icon,
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.addEventListener('mouseenter', Swal.stopTimer)
                toast.addEventListener('mouseleave', Swal.resumeTimer)
            }
        })
        Toast.fire({})
    }

    let success = function (c) {
        const {
            title = "",
            msg = "",
            footer = "",
        } = c;
        Swal.fire({
            icon: 'success',
            title: title,
            text: msg,
            footer: footer,
            color: "white"
        })
    }

    let error = function (c) {
        const {
            title = "",
            msg = "",
            footer = "",
        } = c;
        Swal.fire({
            icon: 'error',
            title: title,
            text: msg,
            footer: footer,
            color: "white"
        })
    }

    let formPopup = async function (c) {
        const {
            msg = "",
            title = ""
        } = c;

        const {value: formValues} = await Swal.fire({
            title: title,
            html: msg,
            focusConfirm: false,
            showCancelButton: true,
            preConfirm: () => {
                return [
                    document.getElementById('start-date').value,
                    document.getElementById('end-date').value
                ]
            },
            willOpen: () => {
                const elem = document.getElementById('reservation-dates-popup');
                const rangePicker = new DateRangePicker(elem, {
                    // Put Options Here
                    format: "dd-mm-yyyy",
                    showFocus: true,
                });
            }
        })

        if (formValues) {
            Swal.fire(JSON.stringify(formValues))
        }
    }

    return {
        toast: toast,
        success: success,
        error: error,
        formPopup: formPopup
    }
}


