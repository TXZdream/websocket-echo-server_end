$(function () {
    var term = new Terminal({
        cursorBlink: true,  // Do not blink the terminal's cursor
        cols: 43,  // Set the terminal's width to 120 columns
        rows: 24  // Set the terminal's height to 80 rows
    });
    term.open(document.getElementById('terminal'));

    var ws = new WebSocket("ws://localhost:8080/echo");

    function runFakeTerminal(term, ws) {
        var str = "";
        var input_state = true;

        if (term._initialized) {
            return;
        }

        term._initialized = true;

        var shellprompt = '$ ';

        term.prompt = function () {
            term.write('\r\n' + shellprompt);
        };

        term.writeln('Welcome to xterm.js');
        term.writeln('Type some keys and commands to play around.');
        term.writeln('');
        term.prompt();

        term.on('key', function (key, ev) {
            var printable = (
                !ev.altKey && !ev.altGraphKey && !ev.ctrlKey && !ev.metaKey
            );

            if (ev.keyCode == 13) {
                term.prompt();
                ws.onopen = function () {
                    ws.send(str);
                    input_state = false;
                }
                ws.onmessage = function (ev) {
                    term.write(ev.data);
                }
                ws.close = function () {
                    str = "";
                    input_state = true;
                }

            } else if (ev.keyCode == 8) {
                // Do not delete the prompt
                if (term.buffer.x > 2) {
                    str = str.slice(0, str.length);
                    term.write('\b \b');
                }
            } else if (printable) {
                if(input_state){
                    str += key;
                }
                term.write(key);
            }
        });

        term.on('paste', function (data, ev) {
            if(input_state){
                str += data;
            }
            term.write(data);
        });
    }
    
});