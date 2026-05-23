import { useState } from 'react';
import api from '../../api/axiosClient';

const SupplierDeliveries = () => {
    const [supplierId, setSupplierId] = useState('');
    const [productId, setProductId] = useState('');
    const [startDate, setStartDate] = useState('');
    const [endDate, setEndDate] = useState('');
    const [deliveries, setDeliveries] = useState([]);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState('');

    const handleSearch = async (e) => {
        e.preventDefault();
        setLoading(true);
        setError('');
        try {
            const url = `/reports/supplier-deliveries?supplier_id=${supplierId}&product_id=${productId}&start_date=${startDate}&end_date=${endDate}`;
            const response = await api.get(url);
            setDeliveries(response.data || []);
        } catch (err) {
            setError('Ошибка загрузки логов поставщика');
        } finally {
            setLoading(false);
        }
    };

    return (
        <div>
            <h2>Поставки определенного товара поставщиком </h2>
            <form onSubmit={handleSearch} style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '10px', marginBottom: '20px', maxWidth: '500px' }}>
                <input type="number" placeholder="ID Поставщика" value={supplierId} onChange={e => setSupplierId(e.target.value)} required style={{ padding: '8px' }}/>
                <input type="number" placeholder="ID Товара" value={productId} onChange={e => setProductId(e.target.value)} required style={{ padding: '8px' }}/>
                <input type="date" placeholder="С даты" value={startDate} onChange={e => setStartDate(e.target.value)} required style={{ padding: '8px' }}/>
                <input type="date" placeholder="По дату" value={endDate} onChange={e => setEndDate(e.target.value)} required style={{ padding: '8px' }}/>
                <button type="submit" style={{ gridColumn: 'span 2', padding: '10px', background: '#1890ff', color: '#fff', border: 'none', cursor: 'pointer' }}>Получить сведения о поставках</button>
            </form>

            {loading && <p>Формирование отчета логистики...</p>}
            {error && <p style={{ color: 'red' }}>{error}</p>}

            <table style={{ width: '100%', borderCollapse: 'collapse' }}>
                <thead>
                <tr style={{ background: '#fafafa', borderBottom: '2px solid #eaeaea' }}>
                    <th style={{ padding: '10px', textAlign: 'left' }}>ID Накладной</th>
                    <th style={{ padding: '10px', textAlign: 'left' }}>Дата Поставки</th>
                    <th style={{ padding: '10px', textAlign: 'left' }}>Поставщик</th>
                    <th style={{ padding: '10px', textAlign: 'left' }}>Товар</th>
                    <th style={{ padding: '10px', textAlign: 'left' }}>Объем (шт)</th>
                    <th style={{ padding: '10px', textAlign: 'left' }}>Закупочная цена</th>
                    <th style={{ padding: '10px', textAlign: 'left' }}>Итоговая стоимость</th>
                </tr>
                </thead>
                <tbody>
                {deliveries.length === 0 ? (
                    <tr><td colSpan="7" style={{ padding: '20px', textAlign: 'center', color: '#999' }}>Данные о поставках за указанный период отсутствуют</td></tr>
                ) : (
                    deliveries.map((d) => (
                        <tr key={d.order_id} style={{ borderBottom: '1px solid #f0f0f0' }}>
                            <td style={{ padding: '10px' }}>{d.order_id}</td>
                            <td style={{ padding: '10px' }}>{new Date(d.order_date).toLocaleDateString()}</td>
                            <td style={{ padding: '10px' }}>{d.supplier_name}</td>
                            <td style={{ padding: '10px' }}>{d.product_name}</td>
                            <td style={{ padding: '10px' }}>{d.quantity}</td>
                            <td style={{ padding: '10px' }}>{d.supply_price} ₽</td>
                            <td style={{ padding: '10px', fontWeight: 'bold' }}>{d.total_cost.toLocaleString()} ₽</td>
                        </tr>
                    ))
                )}
                </tbody>
            </table>
        </div>
    );
};

export default SupplierDeliveries;