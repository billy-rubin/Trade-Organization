import { useState } from 'react';
import api from '../../api/axiosClient';

const CreateRequest = () => {
    const [storeId, setStoreId] = useState('');
    const [productId, setProductId] = useState('');
    const [quantity, setQuantity] = useState(10);
    const [items, setItems] = useState([]);
    const [message, setMessage] = useState('');

    const handleAddItem = (e) => {
        e.preventDefault();
        if (!productId || quantity <= 0) return;
        setItems([...items, { product_id: parseInt(productId), quantity: parseInt(quantity) }]);
        setProductId('');
    };

    const removeItem = (index) => {
        setItems(items.filter((_, i) => i !== index));
    };

    const handleSubmit = async () => {
        if (items.length === 0 || !storeId) {
            setMessage('Укажите ID магазина и добавьте товары!');
            return;
        }

        try {
            const payload = {
                store_id: parseInt(storeId),
                items: items
            };
            const response = await api.post('/supply/requests', payload);
            setMessage(`Заявка #${response.data.request_id} успешно создана!`);
            setItems([]);
            setStoreId('');
        } catch (error) {
            setMessage(`Ошибка: ${error.response?.data?.error || 'Не удалось создать заявку'}`);
        }
    };

    return (
        <div style={{ background: '#fafafa', padding: '20px', borderRadius: '8px' }}>
            <h3>Формирование внутренней заявки</h3>
            {message && <div style={{ padding: '10px', marginBottom: '15px', background: '#e6f7ff', border: '1px solid #91d5ff' }}>{message}</div>}

            <div style={{ marginBottom: '15px' }}>
                <label>ID Магазина (куда требуется товар): </label>
                <input type="number" value={storeId} onChange={(e) => setStoreId(e.target.value)} required />
            </div>

            <form onSubmit={handleAddItem} style={{ display: 'flex', gap: '10px', marginBottom: '20px', alignItems: 'center' }}>
                <input type="number" value={productId} onChange={(e) => setProductId(e.target.value)} placeholder="ID Товара" required />
                <input type="number" value={quantity} onChange={(e) => setQuantity(e.target.value)} min="1" placeholder="Кол-во" required />
                <button type="submit" style={{ padding: '6px 12px' }}>Добавить в список</button>
            </form>

            <ul style={{ listStyle: 'none', padding: 0 }}>
                {items.map((item, index) => (
                    <li key={index} style={{ background: 'white', padding: '10px', border: '1px solid #ddd', marginBottom: '5px', display: 'flex', justifyContent: 'space-between' }}>
                        <span>Товар ID: {item.product_id} | Нужно: {item.quantity} шт.</span>
                        <button onClick={() => removeItem(index)} style={{ color: 'red', border: 'none', background: 'none', cursor: 'pointer' }}>✖</button>
                    </li>
                ))}
            </ul>

            {items.length > 0 && (
                <button onClick={handleSubmit} style={{ marginTop: '15px', padding: '10px 20px', background: '#faad14', color: 'white', border: 'none', borderRadius: '4px', cursor: 'pointer', width: '100%' }}>
                    Отправить заявку
                </button>
            )}
        </div>
    );
};

export default CreateRequest;