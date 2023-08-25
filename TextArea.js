function TextAreaAutoGrow(input) {
    input.style.height = "5px";
    input.style.height = (input.scrollHeight) + "px";
    userTyping(input)
};