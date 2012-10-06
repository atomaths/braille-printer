$(document).ready(function() {
    $("#result").focus(function() {
        $(this).blur();
    });
    //$("#result").autoGrow();
    $("#b-input").focus();
    $("#b-input").keyup(function(e) {
        if (e.keyCode == 13) {
            do_braille();
        }
    });
    $("#btn-braille").click(do_braille);
    $("#btn-print").click(do_print);
});

// 입력한 내용을 서버에 /braille API 호출로 점자로 변환.
function do_braille() {
    var inputval = $.trim($("#b-input").val());
    if (inputval == "") {
        $("#b-input").focus();
        return false;
    }

    $.post("/braille", "b-input=" + inputval, function(result) {
        $("#ly-result").show();
        $("#result").html(result);
        //$("#result").autoGrow();
        $("#b-input").blur();
    });
}

function do_print() {
    alert($("#result").html());
}


/*
// textarea auto grow plugin
jQuery.fn.autoGrow = function(){
    return this.each(function(){
        // Variables
        var colsDefault = this.cols;
        var rowsDefault = this.rows;

        //Functions
        var grow = function() {
            growByRef(this);
        }

        var growByRef = function(obj) {
            var linesCount = 0;
            var lines = obj.value.split('\n');

            for (var i=lines.length-1; i>=0; --i)
    {
        linesCount += Math.floor((lines[i].length / colsDefault) + 1);
    }

    if (linesCount >= rowsDefault)
        obj.rows = linesCount + 1;
    else
        obj.rows = rowsDefault;
        }

        var characterWidth = function (obj){
            var characterWidth = 0;
            var temp1 = 0;
            var temp2 = 0;
            var tempCols = obj.cols;

            obj.cols = 1;
            temp1 = obj.offsetWidth;
            obj.cols = 2;
            temp2 = obj.offsetWidth;
            characterWidth = temp2 - temp1;
            obj.cols = tempCols;

            return characterWidth;
        }

        // Manipulations
        this.style.height = "auto";
        this.style.overflow = "hidden";
        //this.style.width = ((characterWidth(this) * this.cols) + 6) + "px";
        //this.onkeyup = grow;
        //this.onfocus = grow;
        //this.onblur = grow;
        growByRef(this);
    });
};
*/

