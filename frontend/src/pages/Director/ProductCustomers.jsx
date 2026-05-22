import { useState } from 'react';
import api from '../../api/axiosClient';

const ProductCustomers = () => {
    const [productId, setProductId] = useState('');
    const [storeId, setStoreId] = useState('');
    const [customers, setCustomers] = useState([]);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        setLoading(true);
        setError('');
        try {
            const response = await api.get(`/reports/product-customers?product_id=${productId}&store_id=${storeId}`);
            setCustomers(response.data || []);
        } catch (err) {
            setError('Не удалось загрузить лог покупателей');
        } finally {
            setLoading(false);
        }
    };

    return (
        <div>
            <h2>Покупатели конкретного товара (Запрос 10)</h2>
            <form onSubmit={handleSubmit} style={{ display: 'flex', gap: '10px', marginBottom: '20px' }}>
                <input type="number" placeholder="ID Товара" value={productId} onChange={e => setProductId(e.target.value)} required style={{ padding: '8px' }}/>
                <input type="number" placeholder="ID Точки" value={storeId} onChange={e => setStoreId(e.target.value)} required style={{ padding: '8px' }}/>
                <button type="submit" style={{ padding: '8px 16px', background: '#1890ff', color: '#fff', border: 'none', cursor: 'pointer' }}>Найти целевую аудиторию</button>
            </form>

            {loading && <p>Анализ продаж...</p>}
            {error && <p style={{ color: 'red' }}>{error}</p>}

            <table style={{ width: '100%', borderCollapse: 'collapse' }}>
                <thead>
                <tr style={{ background: '#fafafa', borderBottom: '2px solid #eaeaea' }}>
                    <th style={{ padding: '10px', textAlign: 'left' }}>ID Покупателя</th>
                    <th style={{ padding: '10px', textAlign: 'left' }}>ФИО</th>
                    <th style={{ padding: '10px', textAlign: 'left' }}>Характеристики / Примечания</th>
                    <th style={{ padding: '10px', textAlign: 'left' }}>Купленный товар</th>
                </tr>
                </thead>
                <tbody>
                {customers.length === 0 ? (
                    <tr><td colSpan="4" style={{ padding: '20px', textAlign: 'center', color: '#999' }}>Нет записей о покупках данного товара в этой точке</td></tr>
                ) : (
                    customers.map((c) => (
                        <tr key={c.customer_id} style={{ borderBottom: '1px solid #f0f0f0' }}>
                            <td style={{ padding: '10px' }}>{c.customer_id}</td>
                            <td style={{ padding: '10px' }}>{c.customer_name}</td>
                            <td style={{ padding: '10px' }}>{c.characteristics || 'Нет описания'}</td>
                            <td style={{ padding: '10px', color: '#1890ff' }}>{c.bought_product}</td>
                        </tr>
                    ))
                )}
                </tbody>
            </table>
        </div>
    );
};

export default ProductCustomers;