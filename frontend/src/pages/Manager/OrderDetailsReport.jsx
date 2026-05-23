import { useState } from 'react';
import api from '../../api/axiosClient';

const OrderDetailsReport = () => {
    const [orderId, setOrderId] = useState('');
    const [items, setItems] = useState([]);
    const [meta, setMeta] = useState(null);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState('');

    const handleFetchOrder = async (e) => {
        e.preventDefault();
        setLoading(true);
        setError('');
        try {
            const response = await api.get(`/supply/order-details?order_id=${orderId}`);
            setItems(response.data || []);
            if (response.data && response.data.length > 0) {
                setMeta({
                    id: response.data[0].order_id,
                    date: response.data[0].order_date,
                    status: response.data[0].status,
                    supplier: response.data[0].supplier_name
                });
            } else {
                setMeta(null);
            }
        } catch (err) {
            setError('Заказ не найден или у вас нет доступа');
            setItems([]);
            setMeta(null);
        } finally {
            setLoading(false);
        }
    };

    return (
        <div>
            <h2>Просмотр состава и номенклатуры заказа </h2>
            <form onSubmit={handleFetchOrder} style={{ display: 'flex', gap: '10px', marginBottom: '20px' }}>
                <input type="number" placeholder="Введите ID Заказа от поставщика" value={orderId} onChange={e => setOrderId(e.target.value)} required style={{ padding: '8px', width: '250px' }}/>
                <button type="submit" style={{ padding: '8px 16px', background: '#1890ff', color: '#fff', border: 'none', cursor: 'pointer' }}>Открыть спецификацию</button>
            </form>

            {loading && <p>Загрузка номенклатуры...</p>}
            {error && <p style={{ color: 'red' }}>{error}</p>}

            {meta && (
                <div style={{ background: '#f5f5f5', padding: '12px', marginBottom: '15px', borderRadius: '4px', borderLeft: '4px solid #1890ff' }}>
                    <p style={{ margin: '0 0 5px 0' }}><strong>Заказ №:</strong> {meta.id} от {new Date(meta.date).toLocaleDateString()}</p>
                    <p style={{ margin: '0 0 5px 0' }}><strong>Контрагент-поставщик:</strong> {meta.supplier}</p>
                    <p style={{ margin: '0' }}><strong>Текущий статус:</strong> <span style={{ color: meta.status === 'Delivered' ? 'green' : 'orange', fontWeight: 'bold' }}>{meta.status}</span></p>
                </div>
            )}

            <table style={{ width: '100%', borderCollapse: 'collapse' }}>
                <thead>
                <tr style={{ background: '#fafafa', borderBottom: '2px solid #eaeaea' }}>
                    <th style={{ padding: '10px', textAlign: 'left' }}>Наименование товара</th>
                    <th style={{ padding: '10px', textAlign: 'left' }}>Заказанный объем (количество)</th>
                </tr>
                </thead>
                <tbody>
                {items.length === 0 ? (
                    <tr><td colSpan="2" style={{ padding: '20px', textAlign: 'center', color: '#999' }}>Введите корректный номер накладной для отображения спецификации</td></tr>
                ) : (
                    items.map((item, index) => (
                        <tr key={index} style={{ borderBottom: '1px solid #f0f0f0' }}>
                            <td style={{ padding: '10px' }}>{item.product_name}</td>
                            <td style={{ padding: '10px', fontWeight: 'bold' }}>{item.quantity} шт.</td>
                        </tr>
                    ))
                )}
                </tbody>
            </table>
        </div>
    );
};

export default OrderDetailsReport;