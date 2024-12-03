document.addEventListener('DOMContentLoaded', () => {
    const modal = document.getElementById('productModal');
    const addButton = document.getElementById('addProductButton');
    const closeModal = document.getElementById('closeModal');
    const form = document.getElementById('addProductForm');

    // Mostrar el modal
    addButton.addEventListener('click', () => {
        modal.style.display = 'block';
    });

    // Ocultar el modal
    closeModal.addEventListener('click', () => {
        modal.style.display = 'none';
    });

    // Enviar el formulario
    form.addEventListener('submit', (e) => {
        e.preventDefault();
        const product = {
            name: form.name.value,
            price: form.price.value,
            image: form.image.value,
            quantity: form.quantity.value,
        };

        // Llamar a la API para guardar el producto
        fetch('/api/products', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(product),
        })
            .then(res => res.json())
            .then(data => {
                console.log(data);
                modal.style.display = 'none';
                form.reset();
                // Recargar la tabla (puedes implementar esta función)
            })
            .catch(err => console.error(err));
    });
});
