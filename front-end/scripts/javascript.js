
function efeito(menssage,info) {
    
    let nome = info;

    function testando(valor) {
    return menssage + valor;
    }

    console.log("Resultado da função:", testando(nome));
    alert(testando(nome));
}

