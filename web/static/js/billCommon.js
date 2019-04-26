$(document).ready(function() {
  // 顶部按钮
  var divSpan = $(".bill .top div>span:nth-of-type(2)");
  $('.bill .top div span:nth-of-type(1)').on('click',function() {
    if(divSpan.css("display")=="none"){
      divSpan.show();
    }else{
      divSpan.hide();
    }
  });
  // 退出
  $('.exit').on('click',function() {
    $(location).attr('href','/loginout')
  })

  // 折叠菜单
  function filterSpaceNode(node){
		for(var i = 0;i<node.childNodes.length;i++){
				node.removeChild(node.childNodes[i]);
		}
	}

	var box = document.querySelector("#box");
	filterSpaceNode(box);
	var h3 = document.querySelector("#box h3");
	var div = document.querySelector("#box div");
		h3.onclick = function(){
			 none();
			 if(this.nextSibling.style.display=="block"){
				this.nextSibling.style.display="none";
        $('h3 span img').attr('src','../images/jt.png');
			 }else{
				this.nextSibling.style.display = "block";
        $('h3 span img').attr('src','../images/jt2.png');
			 }
	}
	function none(){
			div.style.display.display="none";
	}

  // 发布票据
  $('#list li:nth-of-type(1)').on('click',function() {
    $(location).attr('href','/addBill');
  })
  // 我的票据
  $('#list li:nth-of-type(2)').on('click',function() {
    $(location).attr('href','/bills');
  })
  $('#list li:nth-of-type(3)').on('click',function() {
    $(location).attr('href','/waitEndorseBills');
  })




})
