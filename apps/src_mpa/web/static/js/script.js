function formatRangeDate(index, splitStr) {
  const rows = document.querySelectorAll('tr');

  rows.forEach(row => {
      const dateCell = row.querySelectorAll('td')[index];
      if (dateCell) {
          const rawDates = dateCell.textContent.trim().split(splitStr);
          const startDate = new Date(rawDates[0].trim());
          const endDate = new Date(rawDates[1].trim());
          
          const formattedStartDate = `${startDate.getDate().toString().padStart(2, '0')}.${(startDate.getMonth() + 1).toString().padStart(2, '0')}.${startDate.getFullYear()}`;
          const formattedEndDate = `${endDate.getDate().toString().padStart(2, '0')}.${(endDate.getMonth() + 1).toString().padStart(2, '0')}.${endDate.getFullYear()}`;
          
          dateCell.textContent = `${formattedStartDate}-${formattedEndDate}`;
      }
  });
}

function formatRangeDateTbl(tableID, index, splitStr) {
  var table = document.getElementById(tableID);
  const rows = table.querySelectorAll('tr');

  rows.forEach(row => {
      const dateCell = row.querySelectorAll('td')[index];
      if (dateCell) {
          const rawDates = dateCell.textContent.trim().split(splitStr);
          const startDate = new Date(rawDates[0].trim());
          const endDate = new Date(rawDates[1].trim());
          
          const formattedStartDate = `${startDate.getDate().toString().padStart(2, '0')}.${(startDate.getMonth() + 1).toString().padStart(2, '0')}.${startDate.getFullYear()}`;
          const formattedEndDate = `${endDate.getDate().toString().padStart(2, '0')}.${(endDate.getMonth() + 1).toString().padStart(2, '0')}.${endDate.getFullYear()}`;
          
          dateCell.textContent = `${formattedStartDate}-${formattedEndDate}`;
      }
  });
}

function formatDate(index) {
  const rows = document.querySelectorAll('tr');

  rows.forEach(row => {
      const dateCell = row.querySelectorAll('td')[index];
      if (dateCell) {
          const date = new Date(dateCell.textContent.trim());
          if (!isNaN(date.getTime())) { // Проверка на корректность даты
            const formattedDate = `${date.getDate().toString().padStart(2, '0')}.${(date.getMonth() + 1).toString().padStart(2, '0')}.${date.getFullYear()}`;
            
            dateCell.textContent = `${formattedDate}`;
          }
      }
  });
}

function formatDateTbl(tableID, index) {
  var table = document.getElementById(tableID);
  const rows = table.querySelectorAll('tr');

  rows.forEach(row => {
      const dateCell = row.querySelectorAll('td')[index];
      if (dateCell) {
          const date = new Date(dateCell.textContent.trim());
          
          const formattedDate = `${date.getDate().toString().padStart(2, '0')}.${(date.getMonth() + 1).toString().padStart(2, '0')}.${date.getFullYear()}`;
          
          dateCell.textContent = `${formattedDate}`;
      }
  });
}

function formatBool(value) {
  return value ? 'Да' : 'Нет';
}

function formatBoolCol(){
  const rows = document.querySelectorAll('tr');

  rows.forEach(row => {
      const boolCell = row.querySelector('td.bool-column');
      if (boolCell) {
          const boolValue = boolCell.textContent.trim() === 'true';
          boolCell.textContent = formatBool(boolValue);
      }
  });
}

function formatGenderHelper(value) {
  return value ? 'Женщина' : 'Мужчина';
}

function formatGenderCol(){
  const rows = document.querySelectorAll('tr');

  rows.forEach(row => {
      const boolCell = row.querySelector('td.gender-column');
      if (boolCell) {
          const boolValue = boolCell.textContent.trim() === 'true';
          boolCell.textContent = formatGenderHelper(boolValue);
      }
  });
}

function formatGender() {
  const genderElement = document.querySelector(".gender");
  if (genderElement) {
      const genderValue = genderElement.textContent.trim();
      if (genderValue === "true") {
          genderElement.textContent = "Женщина";
      } else if (genderValue === "false") {
          genderElement.textContent = "Мужчина";
      }
  }
}

function formatMoscowTeam() {
  const genderElement = document.querySelector(".moscowTeam");
  if (genderElement) {
      const genderValue = genderElement.textContent.trim();
      if (genderValue === "true") {
          genderElement.textContent = "Да";
      } else if (genderValue === "false") {
          genderElement.textContent = "Нет";
      }
  }
}

function formatBirthday() {
  const bdayElement = document.querySelector(".bday");

  if (bdayElement) {
      const date = new Date(bdayElement.textContent.trim());
      
      const formattedDate = `${date.getDate().toString().padStart(2, '0')}.${(date.getMonth() + 1).toString().padStart(2, '0')}.${date.getFullYear()}`;
      
      bdayElement.textContent = `${formattedDate}`;
  }
}

function formatDateSpan(className) {
  const bdayElement = document.querySelector(className);

  if (bdayElement) {
      const date = new Date(bdayElement.textContent.trim());
      
      const formattedDate = `${date.getDate().toString().padStart(2, '0')}.${(date.getMonth() + 1).toString().padStart(2, '0')}.${date.getFullYear()}`;
      
      bdayElement.textContent = `${formattedDate}`;
  }
}

function handleRowClick(tableID, resID) {
  var table = document.getElementById(tableID);
  var rows = table.getElementsByTagName('tr');
  
  for (var i = 0; i < rows.length; i++) {
      rows[i].onclick = function() {
      if (this.parentNode.nodeName == 'THEAD') {
          return;
      }
      var cells = this.getElementsByTagName('td');
      var id = cells[0].innerText;
      
      document.getElementById(resID).value = id;

      var selectedRow = table.querySelector('.selected');
      if (selectedRow) {
          selectedRow.classList.remove('selected');
      }

      this.classList.add('selected');
      };
  }
}

function validateDates() {
  document.querySelector('form').addEventListener('submit', function(event) {
    var begDate = new Date(document.getElementById('begDate').value);
    var endDate = new Date(document.getElementById('endDate').value);

    if (endDate < begDate) {
        alert('Дата окончания не может быть раньше даты начала!');
        event.preventDefault(); // Предотвратить отправку формы
    }
    });
}

function validateWeight() {
  document.querySelector('form').addEventListener('submit', function(event) {
    var snatchWeight = parseInt(document.getElementById('snatch').value);
    var cajWeight = parseInt(document.getElementById('caj').value);
    var minWeight = 15;
    if (snatchWeight < minWeight || cajWeight < minWeight) {
        alert('Введенные веса не должны быть меньше 15 кг!');
        event.preventDefault(); // Предотвратить отправку формы
    }
});
}

function validateCompID() {
  document.querySelector('form').addEventListener('submit', function(event) {
    var compIDValue = document.getElementById('compID').value.trim();

    if (compIDValue.length === 0) {
        event.preventDefault(); // Отменяем отправку формы

        alert('Выберите соревнование!'); // Выводим сообщение об ошибке
    }
});
}

function validateSmID() {
  document.querySelector('form').addEventListener('submit', function(event) {
    var compIDValue = document.getElementById('smID').value.trim();

    if (compIDValue.length === 0) {
        event.preventDefault(); // Отменяем отправку формы

        alert('Выберите спортсмена!'); // Выводим сообщение об ошибке
    }
});
}

function validateCampID() {
  document.querySelector('form').addEventListener('submit', function(event) {
    var compIDValue = document.getElementById('campID').value.trim();

    if (compIDValue.length === 0) {
        event.preventDefault(); // Отменяем отправку формы

        alert('Выберите сборы!'); // Выводим сообщение об ошибке
    }
});
}

function validateCID() {
  document.querySelector('form').addEventListener('submit', function(event) {
    var compIDValue = document.getElementById('cID').value.trim();

    if (compIDValue.length === 0) {
        event.preventDefault(); // Отменяем отправку формы

        alert('Выберите тренера!'); // Выводим сообщение об ошибке
    }
});
}