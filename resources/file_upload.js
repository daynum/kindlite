function filePreview(){
  const file = document.querySelector("#highlights").files[0]
  const textSpace = document.querySelector("#highlight-text")
  const reader = new FileReader();

  // I guess we are adding an event listener before reading the file
  // because now when we read the while, the reader gets "loaded"
  // and that event triggers this code which displays the loaded text 
  // on the HTML page.
  reader.addEventListener(
    "load",
    () => {
      console.log(reader.result)
      textSpace.innerHTML = reader.result;
    },
    false,
  );
  console.log(file)
  if(file){
    reader.readAsText(file)
  }
}
