declare namespace Distribute {
    namespace Params {
        interface Query {
            page: number;
            limit: number;
            code?: string;
            status?: number;
        }

        interface Settlement {
            id: number;
            images: string[];
        }
    }
    namespace Response {
       
        type Info = {
            id: number;
            code: string;
            product: {
                id: number;
                name: string;
                pic: string;
                game: string;
                category: string;
            };
            quantity: number;
            unit: string;
            type: number;
            totalAmount: number;
            status: number;
            createTime: string;
        }
        type List = Api.Common.PaginatedResponse<Info>

        type Detail = {
            id: number;
            code: string;
            product: {
                id: number;
                name: string;
                pic: string;
                game: string;
                category: string;
            };
            settlement: {
                id: number;
                code: string;
                amount: number;
                commission: number;
                serviceCharge: number;
                status: number;
                reason: string;
                images: string[];
                createTime: string;
            };
            quantity: number;
            unit: string;
            type: number;
            totalAmount: number;
            status: number;
            requirements: string;
            createTime: string;
            startTime: string;
            finishTime: string;
            reason: string;
        }
    }
}