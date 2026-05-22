import { useState, useEffect } from 'react';
import api from '../../api/axiosClient';

const TurnoverReport = () => {
    const [data, setData] = useState([]);
    const [loading, setLoading] = useState(false);

    const fetchTurnover = async () => {
        setLoading(true);
        try {
            const response = await api.get('/reports/turnover');
            setData(response.data);
        } catch (err) {
            console.error("Ошибка загрузки товарооборота", err);
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => { fetchTurnover(); }, []);

    return (
        <div>
            <h3>Товарооборот по типам торговых точек</h3>
            {loading ? <p>Загрузка...</p> : (
                <table style={{ width: '100%', borderCollapse: 'collapse', marginTop: '15px' }}>
                    <thead>
                    <tr style={{ background: '#fafafa', textAlign: 'left' }}>
                        <th style={{ border: '1px solid #f0f0f0', padding: '8px' }}>Тип точки</th>
                        <th style={{ border: '1px solid #f0f0f0', padding: '8px' }}>Продано товаров (шт.)</th>
                        <th style={{ border: '1px solid #f0f0f0', padding: '8px' }}>Общая выручка (₽)</th>
                    </tr>
                    </thead>
                    <tbody>
                    {data?.map((item, idx) => (
                        <tr key={idx}>
                            <td style={{ border: '1px solid #f0f0f0', padding: '8px' }}>{item.store_type}</td>
                            <td style={{ border: '1px solid #f0f0f0', padding: '8px' }}>{item.total_items_sold}</td>
                            <td style={{ border: '1px solid #f0f0f0', padding: '8px' }}>{item.total_turnover.toLocaleString()} ₽</td>
                        </tr>
                    ))}
                    </tbody>
                </table>
            )}
        </div>
    );
};

export default TurnoverReport;