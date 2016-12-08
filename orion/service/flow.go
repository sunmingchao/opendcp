package service

import (
	"github.com/astaxie/beego/orm"

	"weibo.com/opendcp/orion/models"
	//"weibo.com/opendcp/orion/handler"
)

type FlowService struct {
	BaseService
}

func (f *FlowService) GetFlowWithRel(id int) (*models.Flow, error) {
	o := orm.NewOrm()

	flow := &models.Flow{}
	err := o.QueryTable(flow).Filter("Id", id).RelatedSel().One(flow)
	if err != nil {
		return nil, err
	}

	return flow, nil
}

func (f *FlowService) GetActionImplByName(name string) (*models.ActionImpl, error) {
	o := orm.NewOrm()

	action := &models.ActionImpl{}
	err := o.QueryTable(action).Filter("name", name).One(action)
	if err != nil {
		return nil, err
	}

	return action, nil
}

func (f *FlowService) GetNodeByIp(ip string) (*models.Node, error) {
	o := orm.NewOrm()

	node := &models.Node{}
	err := o.QueryTable(node).Filter("ip", ip).One(node)
	if err != nil {
		return nil, err
	}

	return node, nil
}

/*
func (f *FlowService) GetNodesByFlowId(flowId int) ([]*models.Node, error) {
	o := orm.NewOrm()

	nodeList := make([]*models.Node, 0)

	_, err := o.QueryTable(&models.Node{}).Filter("Flow", flowId).All(&nodeList)

	return nodeList, err
}
*/

func (f *FlowService) GetNodeStatusByFlowId(flowId int) ([]*models.NodeState, error) {
	o := orm.NewOrm()

	nodeList := make([]*models.NodeState, 0)

	_, err := o.QueryTable(&models.NodeState{}).Filter("Flow", flowId).All(&nodeList)

	return nodeList, err
}

/*
func (f *FlowService) GetLog(correlationId string) (string, error) {


}
*/
