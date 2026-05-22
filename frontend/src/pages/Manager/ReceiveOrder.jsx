import { useState } from 'react';
import api from '../../api/axiosClient';

const ReceiveOrder = () => {
    const [orderId, setOrderId] = useState('');
    const [storeId, setStoreId] = useState('');
    const [message, setMessage] = useState('');

    const handleReceive = async (e) => {
        e.preventDefault();
        try {
            await api.post('/supply/orders/receive', {
                order_id: parseInt(orderId),
                store_id: parseInt(storeId)
            });
            setMessage('Товары успешно оприходованы на склад магазина!');
            setOrderId('');
            setStoreId('');
        } catch (error) {
            setMessage(`Ошибка: ${error.response?.data?.error || 'Не удалось принять заказ'}`);
        }
    };

    return (
        <div style={{ background: '#fafafa', padding: '20px', borderRadius: '8px' }}>
            <h3>Приемка товара (Оприходование)</h3>
            <p style={{ fontSize: '14px', color: '#666' }}>Товары из заказа будут добавлены к остаткам магазина.</p>
            {message && <div style={{ padding: '10px', marginBottom: '15px', background: '#f6ffed', border: '1px solid #b7eb8f' }}>{message}</div>}

            <form onSubmit={handleReceive} style={{ display: 'flex', flexDirection: 'column', gap: '15px', maxWidth: '400px' }}>
                <div>
                    <label style={{ display: 'block', marginBottom: '5px' }}>ID Заказа от поставщика:</label>
                    <input type="number" value={orderId} onChange={(e) => setOrderId(e.target.value)} required style={{ width: '100%', padding: '8px' }} />
                </div>
                <div>
                    <label style={{ display: 'block', marginBottom: '5px' }}>ID Магазина (куда привезли):</label>
                    <input type="number" value={storeId} onChange={(e) => setStoreId(e.target.value)} required style={{ width: '100%', padding: '8px' }} />
                </div>

                <button type="submit" style={{ padding: '10px', background: '#52c41a', color: 'white', border: 'none', borderRadius: '4px', cursor: 'pointer' }}>
                    Принять на баланс
                </button>
            </form>
        </div>
    );
};

export default ReceiveOrder;