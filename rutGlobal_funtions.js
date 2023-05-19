
// const input_radio_type_global_run = module.querySelector('input[data-name="global-run"]:checked');
const input_radio_type_global_run_ch = module.querySelector('input[data-name="global-run-ch"]');
const input_radio_type_global_run_ex = module.querySelector('input[data-name="global-run-ex"]');

const input_global_run = module.querySelector('input[data-name="global_run"]');

function changeValueInputRunGlobal(input) {
    let document_type = module.querySelector('input[name="type-global-run"]:checked').value;
    // console.log("CAMBIO EN VALOR RUN GLOBAL: ", input.value,"TIPO: ",document_type);
    if (ValidateGlobalRunDocumentNumber(document_type, input)) {
        reportCorrectInput(input)
    }else{
        reportWrongInput(input)
    }
};

function changeTypeRunGlobal(e) {
    // console.log("CAMBIO TIPO RUN GLOBAL: ", e.target.value);

    // console.log("VALOR RUN ACTUAL: ",input_global_run.value)
    if (ValidateGlobalRunDocumentNumber(e.target.value, input_global_run)){
        reportCorrectInput(input_global_run)
    }else{
        reportWrongInput(input_global_run)
    }
    input_global_run.focus();

};


