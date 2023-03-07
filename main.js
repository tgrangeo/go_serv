console.log("JS loaded")

const url = "localhost:3000"

var todoForm = document.getElementById("newTodo")

todoForm.addEventListener("submit", (e) =>{
	e.preventDefault()

	const formdata = new FormData(todoForm)
	fetch(url,{
		method:"POST",
		body:formData
	}).then(
		response=>response.text()
	).then(
		(data) => {console.log(data) ; document.getElementById("todoList").innerHTML=data}
	).catch(error => console.error(error))

})

function getTodos(){
	fetch("localhost:3000/todos").then((response) =>{
		print(response)
	})
}