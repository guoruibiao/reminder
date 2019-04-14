/*
var items = document.querySelectorAll(".timeline li");
// code for the isElementInViewport function
function callbackFunc() {
  for (var i = 0; i < items.length; i++) {
    if (isElementInViewport(items[i])) {
      items[i].classList.add("in-view");
    }
  }
}
window.addEventListener("load", callbackFunc);
window.addEventListener("scroll", callbackFunc);
*/

function timeformat(timestamp) {
    return new Date(parseInt(timestamp) * 1000).toLocaleString().replace(/:\d{1,2}$/,' '); 
}

function makeNotification(event) {
  var popNotice = function(event) {
    if (Notification.permission == "granted") {
        var notification = new Notification(event.Title, {
            body: event.Description,
            icon: 'https://avatar.csdn.net/0/8/F/3_marksinoberg.jpg'
        });
        
        notification.onclick = function() {
            // TODO just一个展示思路
            // alert(event.Description)
            notification.close();    
        };
    }    
};

  if(Notification.permission == "granted") {
    popNotice(event);   
  }else if(Notification.permission != "denied") {
    Notification.requestPermission(function(permission){
        popNotice(event);
    });
  }
}

function appendEvent(events) {
  for(index=0; index<events.length; index++) {
    var childstr = "<li><div title='"+events[index].Description+"'>"+"<strong><blob><code>"+timeformat(events[index].Tiptime)+"</code></blob></strong><br>"+events[index].Title+"</div><br></li>";
    console.log(childstr);
    $("#bucket").prepend(childstr);
  }
}


function getEvents(all=false) {
  console.log("ready to get events");
  var timestamp = (Date.parse(new Date()))/1000;
    if(all == true) {
      var starttime = 0;
      var endtime = (Date.parse(new Date()))/1000 + 86400 * 90;
    }else{
      var starttime = (Date.parse(new Date()))/1000 - 86400*3;
      var endtime = (Date.parse(new Date()))/1000 + 86400*4;
    }
    $.ajax({
      type:"get",    // 请求类型
      url:"/getevent",    // 请求URL
      data:{starttime:starttime, endtime: endtime},    // 请求参数 即是 在Servlet中 request.getParement();可以得到的字符串
      dataType:"json",    // 数据返回类型
      cache:false, // 是否缓存
      async:true,    // 默认为true 异步请求
      success:function(result){    // 成功返回的结果(响应)
          console.info(result['ret'].length);
          var events = result['ret'];
          if(events && events.length > 0) {
              appendEvent(events);
              // makeNotification(events[0]);
          }else{
              $("#namedlist").html("暂时为空");
          }
      }
  });
}

function getEventForNotification() {
  console.log("ready to get events");
  var timestamp = Date.parse(new Date())/1000;
  var starttime = (Date.parse(new Date()))/1000 - 86400*3;
  var endtime = (Date.parse(new Date()))/1000 + 86400*4;
    $.ajax({
      type:"get",    // 请求类型
      url:"/getevent",    // 请求URL
      data:{starttime:starttime, endtime: endtime},    // 请求参数 即是 在Servlet中 request.getParement();可以得到的字符串
      dataType:"json",    // 数据返回类型
      cache:false, // 是否缓存
      async:true,    // 默认为true 异步请求
      success:function(result){    // 成功返回的结果(响应)
          console.info(result['ret'].length);
          var events = result['ret'];
          if(events && events.length > 0) {
              // 提醒期在一分钟之内的进行notification
              for(index=0; index<events.length; index++) {
                if(events[index].Tiptime > timestamp && events[index].Tiptime <= timestamp + 60) {
                  makeNotification(events[index]);
                }
              }
          }else{
              $("#namedlist").html("暂时为空");
          }
      }
  });
}

function addEvent(title, description) {
  $.ajax({
    type: "get",
    url: "/addevent",
    data: {title: title, description:description},
    dataType: "json",
    cache: false,
    async: true,
    success: function(result) {
      console.log(result);
      $("#btn_addevent").html("添加");
      $("#input_title").attr("value", "");
      $("#input_description").attr("value", "")
    },
    error: function(error) {
      console.log(error)
    }
  });
}

$("#btn_addevent").click(function(){
    $(this).html("隐藏");
    $('#light').css("display", 'block');
    $("#fade").css("display", "block");
});
$("#closeform").click(function(){
  $("#btn_addevent").html("添加");
    $('#light').css("display", 'none');
    $("#fade").css("display", "none");
});

$("#btn_generate").click(function(){
  var title = $("#input_title").val();
  var description = $("#input_description").val();
  console.log(title, description);
  if(title!="" && description.length>11) {
    addEvent(title, description);
    location.reload();
  }else{
    alert("备忘事件内容过短");
  }
});

$(document).ready(function(){
  // 页面加载后加载最新内容
  getEvents(false);
});
setInterval(getEventForNotification, 60*1000);
// setTimeout(makeNotification, 2000);