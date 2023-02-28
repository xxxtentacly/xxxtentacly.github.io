function generateList() {
  var items = [];
  for (var i = 1; i <= 5; i++) {
    items.push({
      index: i,
      text: "Rofl  " + i,
      imagePath: "img" + i + ".jpg"
    });
  }
  return items;
}

function createListItem(item) {
  var li = document.createElement("li");
  li.innerHTML = item.text;
  var img = document.createElement("img");
  img.src = item.imagePath;
  li.appendChild(img);
  return li;
}

function renderList() {
  var list = document.getElementById("list");
  var items = generateList();
  for (var i = 0; i < items.length; i++) {
    var li = createListItem(items[i]);
    list.appendChild(li);
  }
}

renderList();
