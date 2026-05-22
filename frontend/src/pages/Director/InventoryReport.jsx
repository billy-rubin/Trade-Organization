import { useState, useEffect } from 'react';
import api from '../../api/axiosClient';

const InventoryReport = () => {
    const [storeType, setStoreType] = useState('Shop');
    const [data, setData] = useState([]);
    const [loading, setLoading] = useState(false);

    const fetchInventory = async () => {
        setLoading(true);
        try {
            const response = await api.get(`/reports/inventory?store_type=${storeType}`);
            setData(response.data);
        } catch (err) {
            console.error("Ошибка загрузки инвентаря", err);
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => { fetchInventory(); }, [storeType]);

    return (
        <div>
            <h3>Остатки товаров (Инвентаризация)</h3>
            <div style={{ marginBottom: '15px' }}>
                <label>Фильтр по типу точки: </label>
                <select value={storeType} onChange={(e) => setStoreType(e.target.value)} style={{ padding: '5px' }}>
                    <option value="Department_Store">Универмаг</option>
                    <option value="Shop">Магазин</option>
                    <option value="Kiosk">Киоск</option>
                    <option value="Tray">Лоток</option>
                </select>
            </div>

            {loading ? <p>Загрузка...</p> : (
                <table style={{ width: '100%', borderCollapse: 'collapse' }}>
                    <thead>
                    <tr style={{ background: '#fafafa', textAlign: 'left' }}>
                        <th style={{ border: '1px solid #f0f0f0', padding: '8px' }}>ID Точки</th>
                        <th style={{ border: '1px solid #f0f0f0', padding: '8px' }}>Товар</th>
                        <th style={{ border: '1px solid #f0f0f0', padding: '8px' }}>Кол-во на складе</th>
                    </tr>
                    </thead>
                    <tbody>
                    {data?.map((item, idx) => (
                        <tr key={idx}>
                            <td style={{ border: '1px solid #f0f0f0', padding: '8px' }}>{item.store_id}</td>
                            <td style={{ border: '1px solid #f0f0f0', padding: '8px' }}>{item.product_name}</td>
                            <td style={{ border: '1px solid #f0f0f0', padding: '8px' }}>{item.stock_quantity}</td>
                        </tr>
                    ))}
                    </tbody>
                </table>
            )}
        </div>
    );
};

export default InventoryReport;