function TexAreaOninput(input){
    TextAreaAutoGrow(input)
    userFormTyping(input)
}

function TextAreaAutoGrow(input) {
    input.style.height = "5px";
    input.style.height = (input.scrollHeight) + "px";
};




function ResetTextArea(p) {
    
    console.log("ResetTextAreaState:",p)

}