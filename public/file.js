const input = document.querySelector('.fileInput');
const label = input.nextElementSibling;

input.addEventListener('change', e => {
	let fileName = '';
	if (e.target.files && e.target.files.length > 1) {
		fileName = (e.target.getAttribute('data-multiple-caption') || '' ).replace('{count}', e.target.files.length);
	} else {
		fileName = e.target.value.split('\\').pop();
	}
	if (fileName) {
		label.innerHTML = fileName;
	}
});
