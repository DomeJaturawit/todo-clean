package delivery

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	__ "todo-clean/delivery/grpc"
	"todo-clean/domain"
)

//	func (h newHandler) GetTodoDelivery(ctx *gin.Context) {
//		key := ctx.Param("id")
//
//		if key != "" {
//			id, err := uuid.Parse(key)
//			if err != nil {
//				ctx.AbortWithStatusJSON(http.StatusBadRequest, model.GinResponseError{
//					Title: common.ErrFormat.Error(),
//					Error: err.Error(),
//				})
//			}
//			response, err := h.usecase.GetTodoUseCase(ctx, &id)
//			if err != nil {
//				ctx.AbortWithStatusJSON(http.StatusNotFound, model.GinResponseError{
//					Title: common.ErrDataNotFound.Error(),
//					Error: err.Error(),
//				})
//
//			} else {
//				ctx.JSON(http.StatusOK, response)
//			}
//
//		} else {
//			response, err := h.usecase.GetTodoUseCase(ctx, nil)
//			if err != nil {
//				ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.GinResponseError{
//					Title: common.ErrInternal.Error(),
//					Error: err.Error(),
//				})
//
//			}
//			ctx.JSON(http.StatusOK, response)
//		}
//	}
func (n *newHandler) transformTodoRpc(todo []domain.GetTodoEntity) (*[]__.TodoResponse, error) {

	createdAt := &timestamppb.Timestamp{
		Seconds: todo[1].CreatedAt.Unix(),
	}
	uuidString := todo[1].ID.String()

	res := &__.Todo{
		ID:          uuidString,
		Title:       todo[1].Title,
		Description: todo[1].Description,
		Status:      todo[1].Status,
		CreatAt:     createdAt,
	}
	//res := []__.TodoResponse{}
	//for i := range todo {
	//	createdAt := &timestamppb.Timestamp{
	//		Seconds: todo[i].CreatedAt.Unix(),
	//	}
	//	uuidString := todo[i].ID.String()
	//
	//	ress := []__.Todo(
	//		uuidString,
	//		)
	//
	//	resLoop := __.TodoResponse{
	//		Todos: ress,
	//	}
	//
	//	res = append(res, resLoop)
	//
	//}
	//log.Println("res", res)
	return res, nil
}

func (n newHandler) GetTodoDelivery(ctx context.Context, request *__.SingleRequest) (response *[]__.TodoResponse, err error) {
	log.Println("im in", request)
	res, err := n.usecase.GetTodoUseCase(ctx, request.GetId())
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	response, err = n.transformTodoRpc(res)
	log.Println("im out")

	return response, nil
}
