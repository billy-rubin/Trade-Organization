import { useState } from 'react';
import api from '../../api/axiosClient';

const EfficiencyReport = () => {
    const [storeId, setStoreId] = useState('');
    const [data, setData] = useState(null);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState('');

    const handleSearch = async (e) => {
        e.preventDefault();
        if (!storeId) return;

        setLoading(true);
        setError('');
        try {
            // Обращаемся к новому эндпоинту на бэкенде
            const response = await api.get(`/reports/efficiency?store_id=${storeId}`);
            setData(response.data);
        } catch (err) {
            setError('Ошибка загрузки данных или магазин не найден');
            setData(null);
        } finally {
            setLoading(false);
        }
    };

    return (
        <div>
            <h3>Эффективность торговой точки </h3>
            <form onSubmit={handleSearch} style={{ display: 'flex', gap: '10px', marginBottom: '20px' }}>
                <input
                    type="number"
                    placeholder="Введите ID Точки"
                    value={storeId}
                    onChange={e => setStoreId(e.target.value)}
                    required
                    style={{ padding: '5px' }}
                />
                <button type="submit" style={{ padding: '5px 15px', cursor: 'pointer' }}>Показать</button>
            </form>

            {loading && <p>Загрузка...</p>}
            {error && <p style={{ color: 'red' }}>{error}</p>}

            {data && (
                <div style={{ padding: '15px', border: '1px solid #1890ff', borderRadius: '4px', background: '#e6f7ff' }}>
                    <p><strong>ID Точки:</strong> {data.store_id} ({data.store_type})</p>
                    <p><strong>Площадь:</strong> {data.area} кв.м</p>
                    <p><strong>Кол-во прилавков:</strong> {data.counter_count}</p>
                    <p><strong>Общая выручка:</strong> {data.total_revenue.toLocaleString()} ₽</p>
                    <p><strong>Выручка на 1 кв.м:</strong> <span style={{ fontWeight: 'bold', color: '#1890ff' }}>{data.revenue_per_sq_meter.toFixed(2)} ₽/кв.м</span></p>
                </div>
            )}
        </div>
    );
};

export default EfficiencyReport;