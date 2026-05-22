import { useState } from 'react';
import api from '../../api/axiosClient';

const CreateSale = () => {
    const [customerId, setCustomerId] = useState('');
    const [productId, setProductId] = useState('');
    const [quantity, setQuantity] = useState(1);
    const [cart, setCart] = useState([]);
    const [message, setMessage] = useState('');

    // Добавление товара в локальную корзину перед отправкой
    const handleAddToCart = (e) => {
        e.preventDefault();
        if (!productId || quantity <= 0) return;

        setCart([...cart, { product_id: parseInt(productId), quantity: parseInt(quantity) }]);
        setProductId('');
        setQuantity(1);
    };

    // Удаление товара из корзины
    const removeFromCart = (index) => {
        setCart(cart.filter((_, i) => i !== index));
    };

    // Отправка всего чека на сервер
    const handleSubmitSale = async () => {
        if (cart.length === 0) {
            setMessage('Корзина пуста!');
            return;
        }

        try {
            const payload = {
                seller_id: 1,
                customer_id: customerId ? parseInt(customerId) : null,
                items: cart
            };

            const response = await api.post('/trade/sales', payload);
            setMessage(`Успешно! Чек #${response.data.sale_id} оформлен.`);
            setCart([]); // Очищаем корзину после успешной продажи
            setCustomerId('');
        } catch (error) {
            setMessage(`Ошибка: ${error.response?.data?.error || 'Не удалось оформить продажу'}`);
        }
    };

    return (
        <div style={{ background: '#fafafa', padding: '20px', borderRadius: '8px' }}>
            <h3>Новая продажа</h3>
            {message && <div style={{ padding: '10px', marginBottom: '15px', background: '#e6f7ff', border: '1px solid #91d5ff' }}>{message}</div>}

            <div style={{ marginBottom: '15px' }}>
                <label>ID Покупателя (опционально): </label>
                <input type="number" value={customerId} onChange={(e) => setCustomerId(e.target.value)} placeholder="Например: 1" />
            </div>

            <form onSubmit={handleAddToCart} style={{ display: 'flex', gap: '10px', marginBottom: '20px', alignItems: 'center' }}>
                <input type="number" value={productId} onChange={(e) => setProductId(e.target.value)} placeholder="ID Товара" required />
                <input type="number" value={quantity} onChange={(e) => setQuantity(e.target.value)} min="1" placeholder="Кол-во" required />
                <button type="submit" style={{ padding: '6px 12px' }}>Добавить в корзину</button>
            </form>

            <h4>Корзина:</h4>
            <ul style={{ listStyle: 'none', padding: 0 }}>
                {cart.map((item, index) => (
                    <li key={index} style={{ background: 'white', padding: '10px', border: '1px solid #ddd', marginBottom: '5px', display: 'flex', justifyContent: 'space-between' }}>
                        <span>Товар ID: {item.product_id} | Кол-во: {item.quantity} шт.</span>
                        <button onClick={() => removeFromCart(index)} style={{ color: 'red', border: 'none', background: 'none', cursor: 'pointer' }}>✖</button>
                    </li>
                ))}
            </ul>

            {cart.length > 0 && (
                <button onClick={handleSubmitSale} style={{ marginTop: '15px', padding: '10px 20px', background: '#52c41a', color: 'white', border: 'none', borderRadius: '4px', cursor: 'pointer', width: '100%' }}>
                    Пробить чек
                </button>
            )}
        </div>
    );
};

export default CreateSale;