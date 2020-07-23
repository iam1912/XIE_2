$(document).ready(function() {
	$(":submit").click(function() {
		var uname = $("#name").val();
		var upwd = $("#password").val();
		$.ajax({
			type: "post",
			dataType: "text",
			data: $("#form2").serialize(),
			url: "http://localhost:8080/login",
			success: function() {
				window.location.href = "http://localhost:8080/index";
			},
			error: function() {
				window.location.href = "http://localhost:8080/login";
				alert("账号或密码错误,请重新登录!");
			}
		});
	});
});
