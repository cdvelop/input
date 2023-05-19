function TextAreaAutoGrow(input) {
    input.style.height = "5px";
    input.style.height = (input.scrollHeight) + "px";
};

function TextAreaValidate(input) {
    let fd = input.closest('fieldset');
    const rex = /{{.}}/;
    if (rex.test(input.value)) {
       fd.classList.add("foka");
       fd.classList.remove("ferr");
    }else{
       fd.classList.remove("foka");
       fd.classList.add("ferr");
    }
};