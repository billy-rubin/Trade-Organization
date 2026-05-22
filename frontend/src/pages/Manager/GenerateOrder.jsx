import { useState } from 'react';
import api from '../../api/axiosClient';

const GenerateOrder = () => {
    const [requestId, setRequestId] = useState('');
    const [supplierId, setSupplierId] = useState('');
    const [message, setMessage] = useState('');

    const handleGenerate = async (e) => {
        e.preventDefault();
        try {
            await api.post('/supply/orders/generate', {
                request_id: parseInt(requestId),
                supplier_id: parseInt(supplierId)
            });
            setMessage('Заказ успешно сформирован и отправлен поставщику!');
            setRequestId('');
            setSupplierId('');
        } catch (error) {
            setMessage(`Ошибка: ${error.response?.data?.error || 'Не удалось сформировать заказ'}`);
        }
    };

    return (
        <div style={{ background: '#fafafa', padding: '20px', borderRadius: '8px' }}>
            <h3>Сформировать заказ поставщику</h3>
            <p style={{ fontSize: '14px', color: '#666' }}>На основе ранее созданной внутренней заявки.</p>
            {message && <div style={{ padding: '10px', marginBottom: '15px', background: '#e6f7ff', border: '1px solid #91d5ff' }}>{message}</div>}

            <form onSubmit={handleGenerate} style={{ display: 'flex', flexDirection: 'column', gap: '15px', maxWidth: '400px' }}>
                <div>
                    <label style={{ display: 'block', marginBottom: '5px' }}>ID Заявки:</label>
                    <input type="number" value={requestId} onChange={(e) => setRequestId(e.target.value)} required style={{ width: '100%', padding: '8px' }} />
                </div>
                <div>
                    <label style={{ display: 'block', marginBottom: '5px' }}>ID Поставщика:</label>
                    <input type="number" value={supplierId} onChange={(e) => setSupplierId(e.target.value)} required style={{ width: '100%', padding: '8px' }} />
                </div>

                <button type="submit" style={{ padding: '10px', background: '#1890ff', color: 'white', border: 'none', borderRadius: '4px', cursor: 'pointer' }}>
                    Сгенерировать заказ
                </button>
            </form>
        </div>
    );
};

export default GenerateOrder;