function TexAreaOninput(input){
    TextAreaAutoGrow(input)
    userFormTyping(input)
}

function TextAreaAutoGrow(input) {
    input.style.height = "5px";
    input.style.height = (input.scrollHeight) + "px";
};


function ResetTextArea(p) {
    const input = p.form.querySelector('[name="'+p.field_name+'"]')
    TextAreaAutoGrow(input)
    return ""
}