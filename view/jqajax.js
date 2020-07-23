$(document).ready(function() {
	$(":submit").click(function() {
		var uchange = $(this).val();
		var uid = $("#id").val();
		var uname = $("#name").val();
		var umajor = $("#major").val();
		var usocre = $("#socre").val();
		var ubirthday = $("#birthday").val();
		var usex = $("#sex").val();
		$.ajax({
			url: "http://localhost:8080/edit",
			dataType: "json",
			type: "post",
			data: {
				change: uchange,
				id: uid,
				name: uname,
				major: umajor,
				socre: usocre,
				birthday: ubirthday,
				sex: usex
			},
			success: function(data) {
				if (data == "true") {
					alert(uchange + "成功!");
				} else {
					var str = "<tr>" +
						"<th>学号</th>" +
						"<th>姓名</th>" +
						"<th>专业</th>" +
				    	"<th>性别</th>" +
				    	"<th>出生日期</th>" +
				    	"<th>分数</th>" +
						"</tr>"; 
					for(var i=0; i<data.length;i++) {
						str +="<tr>"+
			        	"<td>" + data[i].ID + "</td>" +
						"<td>" + data[i].Name + "</td>" +
						"<td>" + data[i].Major + "</td>" + 
						"<td>" + data[i].Sex + "</td>" +
						"<td>" + data[i].Birthday + "</td>" +
						"<td>" + data[i].Socre + "</td>" +
						"</tr>";
				}
					$("#tab").html(str);
				}
				$(":text").val("");
			},
			error: function() {
				window.location.href = "http://localhost:8080/edit";
				alert(uchange + "失败");
			}
		});
	});
});
