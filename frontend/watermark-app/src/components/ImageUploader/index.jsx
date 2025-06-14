import { useState } from "react";

export default function ImageUploader() {
    const [previews, setPreviews] = useState([])
    const [isDragging, setIsDragging] = useState(false)

    function handleDrop(e) {
        e.preventDefault()
        const files = Array.from(e.dataTransfer.files)
        const urls = files.map(file => ({
            file,
            url: URL.createObjectURL(file)
        }))

        setIsDragging(false)
        setPreviews(urls)
    }

    function handleFileChange() {
        const file = Array.from(e.target.files)
        const urls = files.map(file => ({
            file,
            url: URL.createObjectURL(file),
        }))
        setPreviews(urls)
    }

    function handleOnDragOver(e) {
        e.preventDefault()
        setIsDragging(true)
    }

    function handleOnDragLeave(e) {
        e.preventDefault()
        setIsDragging(false)
    }

    return (
        <>
                {previews.length === 0 ? (
                    <div className='text-center w-full'>
                        <div onDrop={e => handleDrop(e)} onDragOver={handleOnDragOver} onDragLeave={handleOnDragLeave} className={`w-full h-screen flex justify-center items-center flex-col space-y-40 ${ isDragging ? 'bg-blue-100' : '' }`}>
                            <p className="text-gray-500">拖拉圖片到這里，或點擊下方選擇檔案</p>
                            <input
                                type="file"
                                className="file:px-6 file:text-lg file:rounded-full
                                        file:bg-violet-500 file:text-white hover:file:bg-violet-600
                                        cursor-pointer"
                                accept="image/*"
                                onChange={handleFileChange}
                                multiple
                            />
                        </div>
                    </div>
                ) : (
                    // <div className="h-64 overflow-y-auto bg-gray-100 p-4 rounded">
                    //     <div className="grid grid-cols-3 gap-4">
                    //         {previews.map((p, i) => (
                    //         <img
                    //             key={i}
                    //             src={p.url}
                    //             alt={`preview-${i}`}
                    //             className="w-32 h-32 object-cover rounded"
                    //         />
                    //         ))}
                    //     </div>
                    // </div>

                    <div className="h-64 overflow-y-auto bg-gray-100 p-4 rounded">
                        <div className="grid grid-cols-3 gap-4">
                            
                        </div>
                    </div>
                )}
        </>
    )
}