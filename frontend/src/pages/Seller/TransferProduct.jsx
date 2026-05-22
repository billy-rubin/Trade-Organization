import { useState } from 'react';
import api from '../../api/axiosClient';

const TransferProduct = () => {
    const [formData, setFormData] = useState({
        from_store_id: '',
        to_store_id: '',
        product_id: '',
        quantity: ''
    });
    const [message, setMessage] = useState('');

    const handleChange = (e) => {
        setFormData({ ...formData, [e.target.name]: parseInt(e.target.value) || '' });
    };

    const handleTransfer = async (e) => {
        e.preventDefault();
        try {
            await api.post('/trade/transfer', formData);
            setMessage('Товар успешно перемещен!');
            setFormData({ from_store_id: '', to_store_id: '', product_id: '', quantity: '' }); // Сброс формы
        } catch (error) {
            setMessage(`Ошибка: ${error.response?.data?.error || 'Не удалось выполнить перемещение'}`);
        }
    };

    return (
        <div style={{ background: '#fafafa', padding: '20px', borderRadius: '8px' }}>
            <h3>Перемещение товара</h3>
            {message && <div style={{ padding: '10px', marginBottom: '15px', background: '#e6f7ff', border: '1px solid #91d5ff' }}>{message}</div>}

            <form onSubmit={handleTransfer} style={{ display: 'flex', flexDirection: 'column', gap: '15px', maxWidth: '400px' }}>
                <div>
                    <label style={{ display: 'block', marginBottom: '5px' }}>ID Склада-отправителя:</label>
                    <input type="number" name="from_store_id" value={formData.from_store_id} onChange={handleChange} required style={{ width: '100%', padding: '8px' }} />
                </div>
                <div>
                    <label style={{ display: 'block', marginBottom: '5px' }}>ID Склада-получателя:</label>
                    <input type="number" name="to_store_id" value={formData.to_store_id} onChange={handleChange} required style={{ width: '100%', padding: '8px' }} />
                </div>
                <div>
                    <label style={{ display: 'block', marginBottom: '5px' }}>ID Товара:</label>
                    <input type="number" name="product_id" value={formData.product_id} onChange={handleChange} required style={{ width: '100%', padding: '8px' }} />
                </div>
                <div>
                    <label style={{ display: 'block', marginBottom: '5px' }}>Количество:</label>
                    <input type="number" name="quantity" value={formData.quantity} onChange={handleChange} min="1" required style={{ width: '100%', padding: '8px' }} />
                </div>

                <button type="submit" style={{ padding: '10px', background: '#1890ff', color: 'white', border: 'none', borderRadius: '4px', cursor: 'pointer' }}>
                    Оформить перемещение
                </button>
            </form>
        </div>
    );
};

export default TransferProduct;