let colors = ['red', 'orange', 'yellow', 'green', 'blue', 'purple'];
      let i = 0;
      
      setInterval(function() {
        document.getElementById('myDiv').style.backgroundColor = colors[i];
        i++;
        if (i >= colors.length) {
          i = 0;
        }
      }, 1000);