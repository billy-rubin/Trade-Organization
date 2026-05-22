import { useState } from 'react';
import api from '../../api/axiosClient';

const StoreEfficiency = () => {
    const [storeId, setStoreId] = useState('');
    const [data, setData] = useState(null);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState('');

    const handleSearch = async (e) => {
        e.preventDefault();
        setLoading(true);
        setError('');
        try {
            const response = await api.get(`/reports/efficiency?store_id=${storeId}`);
            setData(response.data);
        } catch (err) {
            setError(err.response?.data?.error || 'Ошибка при получении данных');
            setData(null);
        } finally {
            setLoading(false);
        }
    };

    return (
        <div>
            <h2>Эффективность торговой точки (Запрос 7)</h2>
            <form onSubmit={handleSearch} style={{ display: 'flex', gap: '10px', marginBottom: '20px' }}>
                <input
                    type="number"
                    placeholder="ID Торговой точки"
                    value={storeId}
                    onChange={e => setStoreId(e.target.value)}
                    required
                    style={{ padding: '8px', borderRadius: '4px', border: '1px solid #ccc' }}
                />
                <button type="submit" style={{ padding: '8px 16px', background: '#1890ff', color: '#fff', border: 'none', borderRadius: '4px', cursor: 'pointer' }}>Посчитать эффективность</button>
            </form>

            {loading && <p>Расчет данных на сервере...</p>}
            {error && <p style={{ color: 'red' }}>{error}</p>}

            {data && (
                <table style={{ width: '100%', borderCollapse: 'collapse', marginTop: '10px' }}>
                    <thead>
                    <tr style={{ background: '#fafafa', borderBottom: '2px solid #eaeaea' }}>
                        <th style={{ padding: '12px', textAlign: 'left' }}>Параметр</th>
                        <th style={{ padding: '12px', textAlign: 'left' }}>Значение</th>
                    </tr>
                    </thead>
                    <tbody>
                    <tr style={{ borderBottom: '1px solid #f0f0f0' }}><td style={{ padding: '12px' }}>ID Точки / Тип</td><td style={{ padding: '12px' }}>{data.store_id} ({data.store_type})</td></tr>
                    <tr style={{ borderBottom: '1px solid #f0f0f0' }}><td style={{ padding: '12px' }}>Площадь торговой зоны</td><td style={{ padding: '12px' }}>{data.area} кв. м.</td></tr>
                    <tr style={{ borderBottom: '1px solid #f0f0f0' }}><td style={{ padding: '12px' }}>Количество прилавков</td><td style={{ padding: '12px' }}>{data.counter_count} шт.</td></tr>
                    <tr style={{ borderBottom: '1px solid #f0f0f0' }}><td style={{ padding: '12px' }}>Совокупная выручка</td><td style={{ padding: '12px', fontWeight: 'bold', color: '#52c41a' }}>{data.total_revenue.toLocaleString()} ₽</td></tr>
                    <tr style={{ background: '#e6f7ff' }}><td style={{ padding: '12px', fontWeight: 'bold' }}>Эффективность (Выручка / кв.м)</td><td style={{ padding: '12px', fontWeight: 'bold', color: '#1890ff' }}>{data.revenue_per_sq_meter.toFixed(2)} ₽ / кв.м</td></tr>
                    </tbody>
                </table>
            )}
        </div>
    );
};

export default StoreEfficiency;