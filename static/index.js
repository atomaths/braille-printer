$(document).ready(function() {
    set_speech_input_lang();
    var lang = get_browser_lang();
    if (lang != "ko") {
        $("#title-ko").hide();
        $("#title-en").show();
    }

    $("#result").focus(function() {
        $(this).blur();
    });
    //$("#result").autoGrow();
    $("#input").focus();
    $("#input").keyup(function(e) {
        if (e.keyCode == 13) {
            do_braille();
        }
    });
    $("#btn-braille").click(do_braille);
    $("#btn-cloud-print").click(do_print);
    $("#btn-print").click(function() {
        $("#result").print();
    });

    appLogoBaseTop = $(document).height() - 150;
    $("#app-logo").css("margin-top", ($(document).height() - 282) + "px").fadeIn();
});

var appLogoBaseTop = 0;

// 입력한 내용을 서버에 /braille API 호출로 점자로 변환.
function do_braille() {
    $("#app-logo").hide();
    var inputval = $.trim($("#input").val());
    if (inputval == "") {
        $("#input").focus();
        return false;
    }

    var postdata = "input=" + inputval + "&lang=" + get_browser_lang() + "&format=svg";

    $.post("/braille", postdata, function(result) {
        $("#ly-result").show();
        $("#result").html(result);
        //$("#result").autoGrow();
        $("#input").blur();
        
        setTimeout(function() {
            if ($("#btn-print").offset().top > appLogoBaseTop) {
                $("#app-logo").css("margin-top", "25px").fadeIn();
            } else {
                $("#app-logo").css("margin-top", ($(document).height() - (310 + $("#ly-result").outerHeight())) + "px").fadeIn();
            }
        }, 1000);
    });
}

// TODO: 로컬 프린터인지 printq에 넣을지 구분하도록 해야함
function do_print() {
    var inputval = $.trim($("#input").val());
    if (inputval == "") {
        $("#input").focus();
        return false;
    }

    var postdata = "input=" + inputval + "&lang=" + get_browser_lang();

    $.post("/printq/add", postdata, function(result) {
        $("#myModal").modal("show");
        if ($("#ribbon").css("display") == "none") {
            $("#btn-cloud-print").addClass("disabled");
            $("#btn-cloud-print").unbind("click");
            $("#btn-cloud-print").click(function() { alert("일회용 키를 발급받으세요."); });
        }
        /*
        $("#input").blur();
        $("#toast").html("Print reserved.").fadeIn("slow", function() {
            setTimeout(function() {
                $("#toast").fadeOut("slow");
            }, 3000);
        });
        */
    });
}

function set_speech_input_lang() {
    lang = get_browser_lang();
    if (lang == "ko") {
        lang = "ko-KR";
    } else if (lang == "en") {
        lang = "en-US";
    } else {
        lang = "ko-KR";
    }
        
    $("#input").attr("lang", lang);
}

function get_browser_lang() {
    var lang = window.navigator.userLanguage || window.navigator.language;
    lang = lang.toLowerCase().substring(0, 2);
    if (lang == "ko" || lang == "en") {
        return lang;
    } else {
        return "auto";
    }
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

