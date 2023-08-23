function changeDniType(e,form) {
    const input_dni = e.closest('.run-type').querySelector('input[type="text"][data-name="rut_dni"]');

    if (e.value === "ch") {
        input_dni.setAttribute("maxlength", 10);
    } else {
        input_dni.setAttribute("maxlength", 15);
    }
    // console.log("CHANGE DNI TYPE: ", e, " INPUT DNI: ", input_dni);
    UserTyping(input_dni,form);
};

function RunToPointFormat(rut) {
    // XX.XXX.XXX-X
    let run_number = rut.substring(0, rut.length - 2)
    let run_point = FormateaNumero(run_number);
    let _dv = rut.substring(rut.length - 2, rut.length);

    return run_point + _dv
};